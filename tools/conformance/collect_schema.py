#!/usr/bin/env python3
"""Dump a RouterOS menu tree as JSON. Names and structure only. See COLLECTING.md."""

import argparse
import base64
import json
import os
import re
import ssl
import subprocess
import sys
import time
import urllib.error
import urllib.request
from concurrent.futures import ThreadPoolExecutor

HOST = os.environ.get("ROUTEROS_HOST", "")
USER = os.environ.get("ROUTEROS_USER", "admin")
PASSWORD = os.environ.get("ROUTEROS_PASSWORD", "")

MAX_DEPTH = 12
NAME_RE = re.compile(r"^[A-Za-z0-9._-]+$")

# Console verbs and stores: nothing to model, and printing them is slow or huge.
SKIP = {
    "/console", "/environment", "/file", "/log", "/terminal", "/system/history",
    "/system/backup", "/system/upgrade", "/system/package/update", "/tool/sniffer",
}


class Fatal(RuntimeError):
    pass


class Rest:
    """REST transport. Works with a username and password; no SSH key needed."""

    name = "rest"

    def __init__(self, host, user, password, insecure, timeout):
        scheme = "http" if host.startswith("http://") else "https"
        bare = re.sub(r"^https?://", "", host).rstrip("/")
        if not host.startswith(("http://", "https://")):
            scheme = "http"
        self.base = "%s://%s/rest" % (scheme, bare)
        self.timeout = timeout
        self.auth = base64.b64encode(("%s:%s" % (user, password)).encode()).decode()
        self.ctx = None
        if scheme == "https" and insecure:
            self.ctx = ssl.create_default_context()
            self.ctx.check_hostname = False
            self.ctx.verify_mode = ssl.CERT_NONE

    def _call(self, path, body=None, timeout=None):
        req = urllib.request.Request(self.base + path, method="POST" if body else "GET")
        req.add_header("Authorization", "Basic " + self.auth)
        if body is not None:
            req.add_header("Content-Type", "application/json")
            req.data = json.dumps(body).encode()
        with urllib.request.urlopen(req, timeout=timeout or self.timeout,
                                    context=self.ctx) as r:
            return json.loads(r.read().decode() or "null")

    def call(self, path, body=None, timeout=None, tries=3):
        last = None
        for attempt in range(tries):
            try:
                return self._call(path, body, timeout)
            except urllib.error.HTTPError as e:
                if e.code in (401, 403):
                    raise Fatal("%s rejected the credentials for %s (HTTP %d)"
                                % (self.base, USER, e.code))
                return None
            except Exception as e:  # transient: DNS, reset, timeout
                last = e
                time.sleep(0.5 * (attempt + 1))
        raise Fatal("%s is not answering: %s" % (self.base, last))

    def inspect(self, console_path):
        rows = self.call("/console/inspect",
                         {"request": "child", "path": console_path, "as-value": "yes"})
        if not isinstance(rows, list):
            return []
        return [(r.get("name", ""), r.get("node-type", ""), r.get("type", ""))
                for r in rows if isinstance(r, dict)]

    def rows(self, menu):
        got = self.call(menu)
        if isinstance(got, dict):
            return [got]
        return got if isinstance(got, list) else []

    def firmware(self):
        rb = self.call("/system/routerboard")
        return rb.get("current-firmware", "") if isinstance(rb, dict) else ""

    def probe(self):
        r = self.call("/system/resource", timeout=10, tries=1)
        if not isinstance(r, dict) or "board-name" not in r:
            raise Fatal("%s answered, but not like a RouterOS REST API. Is the "
                        "'www' service enabled?" % self.base)
        return r


class Ssh:
    """SSH transport, for a router with the REST service switched off."""

    name = "ssh"

    def __init__(self, host, user, timeout):
        self.host = re.sub(r"^https?://", "", host).rstrip("/")
        self.user = user
        self.timeout = timeout

    def _run(self, script):
        r = subprocess.run(
            ["ssh", "-o", "BatchMode=yes", "-o", "ConnectTimeout=15",
             "%s@%s" % (self.user, self.host), script],
            capture_output=True, text=True, timeout=self.timeout)
        if r.returncode != 0:
            raise Fatal("ssh to %s@%s exited %d: %s"
                        % (self.user, self.host, r.returncode,
                           (r.stderr or r.stdout).strip()[:300] or "no output"))
        return r.stdout

    def inspect(self, console_path):
        out = self._run(':put [:tostr [/console/inspect request=child path="%s" '
                        'as-value]]' % console_path)
        return re.findall(r"name=([^;]+);node-type=([a-z]+);type=([a-z]+)", out)

    def rows(self, menu):
        out = self._run(':do {:put [:tostr [%s/print as-value]]} on-error={:put ""}' % menu)
        rows, cur = [], {}
        for k in re.findall(r"([A-Za-z0-9.\-]+)=", out):
            if k == ".id" and cur:
                rows.append(cur)
                cur = {}
            cur[k] = ""
        if cur:
            rows.append(cur)
        return rows

    def firmware(self):
        try:
            return self._run(':do {:put [/system/routerboard/get current-firmware]} '
                             'on-error={:put ""}').strip()
        except Fatal:
            return ""

    def probe(self):
        out = self._run(':put [/system/resource/get board-name]; '
                        ':put [/system/resource/get version]; '
                        ':put [/system/resource/get architecture-name]')
        v = [l.strip() for l in out.splitlines() if l.strip()]
        if not v:
            raise Fatal("ssh connected but /system/resource returned nothing")
        v += ["", "", ""]
        return {"board-name": v[0], "version": v[1], "architecture-name": v[2]}


def console_path(menu, *rest):
    return ",".join([p for p in menu.strip("/").split("/") if p] + list(rest))


def progress(msg):
    sys.stderr.write(msg + "\n")
    sys.stderr.flush()


def walk(dev, root, jobs):
    """{menu: node type}. `path` is a namespace (/ip); `dir` is a real menu."""
    found, frontier, depth, truncated = {}, {root: "path"}, 0, False
    while frontier:
        if depth >= MAX_DEPTH:
            truncated = True
            break
        with ThreadPoolExecutor(max_workers=jobs) as pool:
            listings = dict(zip(sorted(frontier),
                                pool.map(lambda m: dev.inspect(console_path(m)),
                                         sorted(frontier))))
        nxt = {}
        for menu, kind in frontier.items():
            if menu:
                found[menu] = kind
            for name, node_type, rel in listings.get(menu, []):
                if rel != "child" or node_type not in ("dir", "path"):
                    continue
                child = ("%s/%s" % (menu, name)).replace("//", "/")
                if child in SKIP or any(child.startswith(s + "/") for s in SKIP):
                    continue
                if child in found:
                    continue
                nxt[child] = node_type
        progress("  depth %d: %d menus below" % (depth, len(nxt)))
        frontier, depth = nxt, depth + 1
    if truncated:
        progress("  WARNING: stopped at depth %d; the tree may be deeper" % MAX_DEPTH)
    return found, truncated


def args_for(dev, menus, verb, jobs):
    def one(m):
        return [n for n, node_type, rel in dev.inspect(console_path(m, verb))
                if node_type == "arg" and rel == "child" and n != "numbers"]
    with ThreadPoolExecutor(max_workers=jobs) as pool:
        return dict(zip(menus, pool.map(one, menus)))


def row_shape(dev, menus, jobs):
    def one(m):
        try:
            rows = dev.rows(m)
        except Fatal:
            raise
        except Exception:
            return (0, [])
        keys = set()
        for r in rows:
            if isinstance(r, dict):
                keys.update(r.keys())
        # Bucketed, not counted: how many firewall connections or bridge hosts
        # someone has describes their network, and none / one / many is all the
        # schema needs to tell a singleton from a collection.
        n = len(rows)
        return ("none" if n == 0 else "one" if n == 1 else "many", sorted(keys))
    with ThreadPoolExecutor(max_workers=jobs) as pool:
        return dict(zip(menus, pool.map(one, menus)))


def audit(doc):
    """Refuse to emit anything that is not a bare name."""
    allowed = {"collector_version", "contains", "identity", "root", "truncated", "menus"}
    extra = set(doc) - allowed
    if extra:
        raise Fatal("unexpected top-level keys: %s" % sorted(extra))
    for menu, e in doc["menus"].items():
        if e["node_type"] not in ("dir", "path"):
            raise Fatal("%s has node_type %r" % (menu, e["node_type"]))
        if e["rows"] not in ("none", "one", "many"):
            raise Fatal("%s has rows %r" % (menu, e["rows"]))
        for field in ("set_args", "add_args", "commands", "submenus", "row_keys"):
            for name in e[field]:
                if not NAME_RE.match(name):
                    raise Fatal("%s.%s holds %r, which is not a bare name"
                                % (menu, field, name))
    return doc


def collect(dev, root, with_rows, jobs):
    progress("checking the connection")
    res = dev.probe()
    ident = {
        "board": res.get("board-name", ""),
        "version": res.get("version", ""),
        "architecture": res.get("architecture-name", ""),
        "firmware": "",
        "transport": dev.name,
    }
    try:
        ident["firmware"] = dev.firmware()
    except Exception:
        pass
    progress("  %s / %s / %s" % (ident["board"], ident["version"], ident["architecture"]))

    if root and not dev.inspect(console_path(root)):
        raise Fatal("%s is not a menu on this device" % root)

    progress("walking the menu tree")
    kinds, truncated = walk(dev, root, jobs)
    menus = sorted(kinds)
    progress("  %d menus" % len(menus))
    if not menus:
        raise Fatal("no menus found below %r" % (root or "/"))

    progress("set arguments");   set_args = args_for(dev, menus, "set", jobs)
    progress("add arguments");   add_args = args_for(dev, menus, "add", jobs)
    progress("menu children")
    with ThreadPoolExecutor(max_workers=jobs) as pool:
        children = dict(zip(menus, pool.map(lambda m: dev.inspect(console_path(m)), menus)))

    rows = {}
    if with_rows:
        progress("row shapes")
        rows = row_shape(dev, [m for m in menus if kinds[m] == "dir"], jobs)

    out = {}
    for m in menus:
        kids = children.get(m, [])
        bucket, keys = rows.get(m, ("none", []))
        e = {
            "node_type": kinds[m],
            "set_args": sorted(set(set_args.get(m) or [])),
            "add_args": sorted(set(add_args.get(m) or [])),
            "commands": sorted({n for n, t, rel in kids if t == "cmd" and rel == "child"}),
            "submenus": sorted({n for n, t, rel in kids
                                if t in ("dir", "path") and rel == "child"}),
            "rows": bucket,
            "row_keys": keys,
        }
        if any(e[k] for k in ("set_args", "add_args", "commands", "submenus", "row_keys")):
            out[m] = e

    if not out:
        raise Fatal("walked %d menus below %r but none had any arguments, "
                    "commands or properties" % (len(menus), root or "/"))

    return audit({
        "collector_version": 2,
        "contains": "names and structure only; no values, no counts",
        "identity": ident,
        "root": root or "/",
        "truncated": truncated,
        "menus": out,
    })


def summarise(doc):
    i = doc["identity"]
    print("%s  %s  %s  firmware %s  (via %s)"
          % (i["board"], i["version"], i["architecture"], i["firmware"] or "?", i["transport"]))
    m = doc["menus"]
    print("%d menus, %d settable arguments, %d row properties"
          % (len(m), sum(len(e["set_args"]) for e in m.values()),
             sum(len(e["row_keys"]) for e in m.values())))
    print("%d menus hold rows" % len([1 for e in m.values() if e["rows"] != "none"]))
    if doc["truncated"]:
        print("WARNING: the walk hit its depth limit; some menus may be missing")


def main():
    ap = argparse.ArgumentParser(description="Dump a RouterOS menu tree. Names only.")
    ap.add_argument("--root", default="", help="subtree to walk, e.g. /interface/ethernet")
    ap.add_argument("--transport", choices=("rest", "ssh"), default="rest")
    ap.add_argument("--no-rows", action="store_true", help="skip reading rows entirely")
    ap.add_argument("--check", action="store_true", help="print a summary instead of JSON")
    ap.add_argument("--insecure", action="store_true", help="accept a self-signed TLS cert")
    ap.add_argument("--jobs", type=int, default=4, help="concurrent requests (default 4)")
    ap.add_argument("--timeout", type=int, default=60, help="per-request seconds")
    ap.add_argument("-o", "--out", help="write JSON here instead of stdout")
    a = ap.parse_args()

    if not HOST:
        sys.exit("ROUTEROS_HOST is required, e.g. ROUTEROS_HOST=192.168.88.1")
    if a.transport == "rest" and not PASSWORD:
        sys.exit("ROUTEROS_PASSWORD is required for the REST transport "
                 "(or use --transport ssh with an SSH key)")

    dev = (Rest(HOST, USER, PASSWORD, a.insecure, a.timeout)
           if a.transport == "rest" else Ssh(HOST, USER, a.timeout))

    doc = collect(dev, a.root.rstrip("/"), not a.no_rows, max(1, a.jobs))
    if a.check:
        summarise(doc)
    else:
        text = json.dumps(doc, indent=1, sort_keys=True) + "\n"
        if a.out:
            with open(a.out, "w") as fh:
                fh.write(text)
            progress("wrote %s (%d KB)" % (a.out, len(text) // 1024))
        else:
            sys.stdout.write(text)
    progress("done")


if __name__ == "__main__":
    try:
        main()
    except Fatal as e:
        sys.exit("collection aborted: %s" % e)
    except KeyboardInterrupt:
        sys.exit("interrupted")
