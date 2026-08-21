#!/usr/bin/env python3
"""Dump a RouterOS menu tree as JSON. Names only, never values. See COLLECTING.md."""

import argparse
import json
import os
import re
import subprocess
import sys

HOST = os.environ.get("ROUTEROS_HOST", "").replace("http://", "").replace("https://", "")
USER = os.environ.get("ROUTEROS_USER", "admin")

BATCH = 40
MAX_DEPTH = 8

# Console verbs and log/file stores: nothing to model, slow or noisy to print.
SKIP_MENUS = {
    "/console", "/environment", "/file", "/log", "/terminal", "/system/history",
    "/system/backup", "/system/upgrade", "/system/package/update",
}


class SSHError(RuntimeError):
    pass


def ssh(script, timeout=240):
    r = subprocess.run(
        ["ssh", "-o", "BatchMode=yes", "-o", "ConnectTimeout=15",
         "%s@%s" % (USER, HOST), script],
        capture_output=True, text=True, timeout=timeout)
    if r.returncode != 0:
        raise SSHError("ssh to %s@%s exited %d: %s"
                       % (USER, HOST, r.returncode,
                          (r.stderr or r.stdout).strip()[:400] or "no output"))
    return r.stdout


def console_path(menu, *rest):
    return ",".join([p for p in menu.strip("/").split("/") if p] + list(rest))


def progress(msg):
    sys.stderr.write(msg + "\n")
    sys.stderr.flush()


def _batched(items, render, parse):
    out = {}
    total = (len(items) + BATCH - 1) // BATCH
    for i in range(0, len(items), BATCH):
        cmds = []
        for it in items[i:i + BATCH]:
            cmds.append(':put "@@@%s"' % (it if isinstance(it, str) else "%s|%s" % it))
            cmds.append(render(it))
        cur = None
        for line in ssh("; ".join(cmds)).splitlines():
            line = line.strip()
            if line.startswith("@@@"):
                cur = line[3:]
                out.setdefault(cur, parse(None, None))
            elif cur is not None:
                out[cur] = parse(out[cur], line)
        progress("  %d/%d" % (i // BATCH + 1, total))
    return out


def inspect_children(menus):
    return _batched(
        menus,
        lambda m: ':put [:tostr [/console/inspect request=child path="%s" as-value]]'
                  % console_path(m),
        lambda acc, line: [] if line is None
        else acc + re.findall(r"name=([^;]+);node-type=([a-z]+);type=child", line))


def inspect_args(menus, verb):
    return _batched(
        menus,
        lambda m: ':put [:tostr [/console/inspect request=child path="%s" as-value]]'
                  % console_path(m, verb),
        lambda acc, line: [] if line is None
        else acc + [a for a in re.findall(r"name=([^;]+);node-type=arg", line)
                    if a != "numbers"])


def row_shape(menus):
    raw = _batched(
        menus,
        lambda m: ':do {:put [:tostr [%s/print as-value]]} on-error={:put ""}' % m,
        lambda acc, line: (0, set()) if line is None
        else (max(acc[0], line.count(".id=") or (1 if "=" in line else 0)),
              acc[1] | set(re.findall(r"([A-Za-z0-9.\-]+)=", line))))
    return {m: (c, sorted(k)) for m, (c, k) in raw.items()}


def enum_values(pairs):
    raw = _batched(
        pairs,
        lambda p: ':put [:tostr [/console/inspect request=completion path=%s as-value]]'
                  % console_path(p[0], "set", p[1]),
        lambda acc, line: [] if line is None
        else [v for v, show in re.findall(
            r"completion=(.*?);offset=\d+;preference=-?\d+;show=(\w+);", line)
            if show == "true"] or acc)
    return {k.replace("|", " "): v for k, v in raw.items() if v}


def walk(root):
    """{menu: node kind}. `path` is a namespace (/ip); `dir` is a real menu."""
    found, frontier, depth = {}, {root: "path"}, 0
    while frontier and depth < MAX_DEPTH:
        children = inspect_children(sorted(frontier))
        nxt = {}
        for menu, kind in frontier.items():
            if menu:
                found[menu] = kind
            for name, child_kind in children.get(menu, []):
                if child_kind not in ("dir", "path"):
                    continue
                child = ("%s/%s" % (menu, name)).replace("//", "/")
                if child in SKIP_MENUS or any(child.startswith(s + "/") for s in SKIP_MENUS):
                    continue
                nxt[child] = child_kind
        progress("depth %d: %d menus below" % (depth, len(nxt)))
        frontier, depth = nxt, depth + 1
    return found


def identity():
    out = ssh(':put [/system/resource/get board-name]; '
              ':put [/system/resource/get version]; '
              ':put [/system/resource/get architecture-name]; '
              ':do {:put [/system/routerboard/get current-firmware]} on-error={:put ""}')
    vals = [l.strip() for l in out.splitlines()]
    keys = ["board", "version", "architecture", "firmware"]
    return dict(zip(keys, vals + [""] * len(keys)))


def audit_output(doc):
    """Refuse to emit anything that looks like configuration rather than schema."""
    allowed = {"collector_version", "contains", "identity", "root", "menus", "enums"}
    extra = set(doc) - allowed
    if extra:
        raise ValueError("unexpected top-level keys: %s" % sorted(extra))
    for menu, e in doc["menus"].items():
        if e["node_type"] not in ("dir", "path"):
            raise ValueError("%s has node_type %r" % (menu, e["node_type"]))
        for field in ("set_args", "add_args", "commands", "submenus", "row_keys"):
            for name in e[field]:
                if not re.fullmatch(r"[A-Za-z0-9._\-]+", name):
                    raise ValueError("%s.%s holds %r, which is not a bare name"
                                     % (menu, field, name))
    return doc


def collect(root, with_enums, with_rows=True):
    progress("identity")
    ident = identity()
    progress("  %s / %s / %s" % (ident["board"], ident["version"], ident["architecture"]))

    progress("walking the menu tree")
    kinds = walk(root)
    menus = sorted(kinds)
    progress("  %d menus" % len(menus))

    progress("set arguments");  set_args = inspect_args(menus, "set")
    progress("add arguments");  add_args = inspect_args(menus, "add")
    progress("menu children");  children = inspect_children(menus)
    rows = {}
    if with_rows:
        progress("row shapes")
        rows = row_shape([m for m in menus if kinds[m] == "dir"])

    out = {}
    for m in menus:
        kids = children.get(m, [])
        count, keys = rows.get(m, (0, []))
        e = {
            "node_type": kinds[m],
            "set_args": sorted(set(set_args.get(m, []))),
            "add_args": sorted(set(add_args.get(m, []))),
            "commands": sorted({n for n, k in kids if k == "cmd"}),
            "submenus": sorted({n for n, k in kids if k in ("dir", "path")}),
            "row_count": count,
            "row_keys": keys,
        }
        if any(e[k] for k in ("set_args", "add_args", "commands", "submenus", "row_keys")):
            out[m] = e

    doc = {
        "collector_version": 1,
        "contains": "names only; no configuration values",
        "identity": ident,
        "root": root or "/",
        "menus": out,
    }
    if with_enums:
        progress("enum value sets")
        doc["enums"] = enum_values([(m, a) for m, e in out.items() for a in e["set_args"]])
    return audit_output(doc)


def summarise(doc):
    i = doc["identity"]
    print("%s  %s  %s  firmware %s" % (i["board"], i["version"], i["architecture"], i["firmware"]))
    menus = doc["menus"]
    print("%d menus, %d settable arguments, %d row properties"
          % (len(menus), sum(len(e["set_args"]) for e in menus.values()),
             sum(len(e["row_keys"]) for e in menus.values())))
    print("%d menus have at least one row"
          % len([1 for e in menus.values() if e["row_count"]]))
    if "enums" in doc:
        print("%d enum value sets" % len(doc["enums"]))


def main():
    ap = argparse.ArgumentParser(description="Dump a RouterOS menu tree. Names only.")
    ap.add_argument("--root", default="", help="subtree to walk, e.g. /interface/ethernet")
    ap.add_argument("--with-enums", action="store_true", help="also record enum value sets")
    ap.add_argument("--no-rows", action="store_true",
                    help="skip reading rows; loses read-only property names")
    ap.add_argument("--check", action="store_true", help="print a summary instead of JSON")
    a = ap.parse_args()
    if not HOST:
        sys.exit("ROUTEROS_HOST is required, e.g. ROUTEROS_HOST=192.168.88.1")
    doc = collect(a.root.rstrip("/"), a.with_enums, not a.no_rows)
    if a.check:
        summarise(doc)
    else:
        json.dump(doc, sys.stdout, indent=1, sort_keys=True)
        sys.stdout.write("\n")
    progress("done")


if __name__ == "__main__":
    try:
        main()
    except SSHError as e:
        sys.exit("collection aborted: %s" % e)
