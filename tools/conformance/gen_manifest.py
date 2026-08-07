#!/usr/bin/env python3
"""Build a conformance manifest: every resource, every attribute, real values.

The acceptance suite was green through 202 broken property names because 296 of
its 364 files set no attribute at all. This closes that hole by generating, from
the device's own schema, a manifest that drives every settable attribute of every
reachable resource through create -> update -> destroy with two *different*
valid values, so both the write path and the read-back path are exercised.

Values come from the router, not from guesses:
  * enums          `/console/inspect request=completion` lists the accepted
                   tokens; two distinct ones are picked.
  * free values    the provider's own validator says what shape is legal
                   (IsIP, IsCIDR, IsMAC, IsDurationRouterOS, OneOf), so the
                   generated value matches the attribute's semantic type.
  * bool/int       true/false, and two small integers.

Attributes and whole resources that can sever management access are excluded by
name -- addresses, services, users, firewall filters and friends. The dead-man
switch in deadman.sh is the backstop, not the first line of defence.

Usage:
    ROUTEROS_HOST=... ROUTEROS_USER=... ROUTEROS_PASSWORD=... \
        python3 tools/conformance/gen_manifest.py > \
        internal/provider/testdata/conformance.json
"""

import json
import os
import re
import subprocess
import sys
from glob import glob

HOST = os.environ.get("ROUTEROS_HOST", "").replace("http://", "").replace("https://", "")
USER = os.environ.get("ROUTEROS_USER", "admin")
SRC = "internal/provider"
BATCH = 20

# Touching these can cut the management path. The dead-man switch would recover
# the router, but a sweep that trips it every few seconds makes no progress.
UNSAFE_MENUS = {
    "/ip/address", "/ip/service", "/ip/firewall/filter", "/ipv6/firewall/filter",
    "/ip/firewall/nat", "/user", "/user/group", "/interface/bridge",
    "/interface/bridge/port", "/interface/bridge/vlan", "/interface/ethernet",
    "/system/identity", "/system/routerboard/settings", "/ip/route",
    "/ip/dhcp-client", "/ip/dhcp-server", "/tool/mac-server", "/ip/neighbor/discovery-settings",
    "/interface/list", "/interface/list/member", "/ip/cloud", "/system/scheduler",
    "/system/script", "/ip/ssh", "/ip/settings", "/system/package/update",
}
UNSAFE_ATTRS = {
    "disabled", "name", "address", "interface", "port", "src_address",
    "dst_address", "chain", "action", "password", "user", "group", "policy",
    "mac_address", "vlan_id", "bridge", "master_interface", "certificate",
}


def ssh(script, timeout=240):
    return subprocess.run(
        ["ssh", "-o", "BatchMode=yes", "-o", "ConnectTimeout=15",
         "%s@%s" % (USER, HOST), script],
        capture_output=True, text=True, timeout=timeout).stdout


def cpath(menu, *rest):
    return ",".join([menu.strip("/").replace("/", ",")] + list(rest))


def provider_resources():
    """resource type name -> {menu, file, attrs: {tfsdk: {kind, validator}}}."""
    out = {}
    for f in sorted(glob(SRC + "/resource_*.go")):
        if f.endswith("_test.go"):
            continue
        src = open(f).read()
        tn = re.search(r'ProviderTypeName \+ "(_[a-z0-9_]+)"', src)
        menus = sorted(set(re.findall(r'ctx,\s*"(/[a-z0-9/\-]+)"', src)))
        if not tn or len(menus) != 1:
            continue
        attrs = {}
        for m in re.finditer(
                r'"([a-z0-9_]+)":\s*schema\.(\w+?)Attribute\{(.{0,600}?)(?=\n\t{3}"[a-z0-9_]+":|\n\t{2}\},)',
                src, re.S):
            name, kind, body = m.group(1), m.group(2), m.group(3)
            if "Optional:" not in body:
                continue                      # Computed-only: not settable
            vld = re.search(r'schemautil\.(\w+)\(([^)]*)\)', body)
            enum = re.findall(r'"([^"]+)"', vld.group(2)) if vld and vld.group(1).startswith("OneOf") else []
            attrs[name] = {"kind": kind, "validator": vld.group(1) if vld else "",
                           "enum": enum}
        out["routeros" + tn.group(1)] = {"menu": menus[0],
                                         "file": os.path.basename(f),
                                         "attrs": attrs}
    return out


def device_menu_info(menus):
    """menu -> {"args": [...], "addable": bool, "rows": int}."""
    info = {}
    for i in range(0, len(menus), BATCH):
        cmds = []
        for m in menus[i:i + BATCH]:
            cmds.append(':put "@@@%s"' % m)
            cmds.append(':put [:tostr [/console/inspect request=child path=%s as-value]]'
                        % cpath(m, "set"))
            cmds.append(':put [:tostr [/console/inspect request=child path=%s as-value]]'
                        % cpath(m))
        cur, phase = None, 0
        for line in ssh("; ".join(cmds)).splitlines():
            line = line.strip()
            if line.startswith("@@@"):
                cur = line[3:]; phase = 0
                info.setdefault(cur, {"args": [], "addable": False})
            elif cur and "name=" in line:
                phase += 1
                if phase == 1:
                    info[cur]["args"] = re.findall(r"name=([^;]+);node-type=arg", line)
                else:
                    info[cur]["addable"] = "name=add;node-type=cmd" in line
        sys.stderr.write("  menus %d/%d\n" % (i // BATCH + 1, (len(menus) + BATCH - 1) // BATCH))
    return info


def device_values(pairs):
    """(menu, arg) -> list of accepted tokens ([] when it takes a free value)."""
    res = {}
    for i in range(0, len(pairs), BATCH):
        cmds = []
        for menu, arg in pairs[i:i + BATCH]:
            cmds.append(':put "@@@%s|%s"' % (menu, arg))
            cmds.append(':put [:tostr [/console/inspect request=completion path=%s as-value]]'
                        % cpath(menu, "set", arg))
        cur = None
        for line in ssh("; ".join(cmds)).splitlines():
            line = line.strip()
            if line.startswith("@@@"):
                cur = tuple(line[3:].split("|", 1)); res.setdefault(cur, [])
            elif cur and "completion=" in line:
                res[cur] = [v for v, _, show, _ in re.findall(
                    r"completion=(.*?);offset=\d+;preference=(-?\d+);show=(\w+);style=(\S*?);",
                    line) if show == "true" and v != "<value>"]
        sys.stderr.write("  values %d/%d\n" % (i // BATCH + 1, (len(pairs) + BATCH - 1) // BATCH))
    return res


FREE = {
    "IsIP":              ("192.0.2.10", "192.0.2.20"),
    "IsCIDR":            ("192.0.2.0/24", "198.51.100.0/24"),
    "IsMAC":             ("02:00:00:00:00:01", "02:00:00:00:00:02"),
    "IsDurationRouterOS": ("30s", "1m"),
    "IsDurationOrKeyword": ("30s", "1m"),
}


def value_pair(attr, tokens):
    """Two different legal values for one attribute, or None to skip it."""
    if tokens:
        toks = [t for t in tokens if t not in ("*",)]
        if len(toks) >= 2:
            return toks[0], toks[1]
        return None
    if attr["enum"] and len(attr["enum"]) >= 2:
        return attr["enum"][0], attr["enum"][1]
    if attr["validator"] in FREE:
        return FREE[attr["validator"]]
    if attr["kind"] == "Bool":
        return "true", "false"
    if attr["kind"] == "Int64":
        return "10", "20"
    if attr["kind"] == "String" and not attr["validator"]:
        return "tf-acc-a", "tf-acc-b"
    return None


def main():
    if not HOST:
        sys.exit("ROUTEROS_HOST is required")
    here = os.path.dirname(os.path.abspath(__file__))
    with open(os.path.join(here, "overrides.json")) as fh:
        overrides = {k: v for k, v in json.load(fh).items() if not k.startswith("_")}
    resources = provider_resources()
    menus = sorted({r["menu"] for r in resources.values()})
    sys.stderr.write("resources: %d across %d menus\n" % (len(resources), len(menus)))
    info = device_menu_info(menus)

    wanted = []
    for res in resources.values():
        args = info.get(res["menu"], {}).get("args", [])
        for a in args:
            if a in ("numbers", "comment"):
                continue
            wanted.append((res["menu"], a))
    wanted = sorted(set(wanted))
    sys.stderr.write("attributes to price: %d\n" % len(wanted))
    vals = device_values(wanted)

    manifest, skipped = {}, {}
    for name, res in sorted(resources.items()):
        menu = res["menu"]
        mi = info.get(menu, {})
        if not mi.get("args"):
            skipped[name] = "menu absent on this device"
            continue
        if menu in UNSAFE_MENUS:
            skipped[name] = "excluded: management path"
            continue
        cases = {}
        for arg in mi["args"]:
            if arg in ("numbers",):
                continue
            tf = arg.replace("-", "_").replace(".", "_")
            if tf not in res["attrs"] or tf in UNSAFE_ATTRS:
                continue
            pair = value_pair(res["attrs"][tf], vals.get((menu, arg), []))
            ov = overrides.get("%s.%s" % (name, tf))
            if ov:
                pair = (ov["a"], ov["b"])
            if pair:
                cases[tf] = {"kind": res["attrs"][tf]["kind"],
                             "a": pair[0], "b": pair[1]}
        if not cases:
            skipped[name] = "no safely settable attribute"
            continue
        manifest[name] = {"menu": menu, "addable": mi.get("addable", False),
                          "key": "name" if "name" in mi["args"] else "",
                          "attributes": cases}

    covered = sum(len(v["attributes"]) for v in manifest.values())
    sys.stderr.write("MANIFEST: %d resources, %d attributes, %d skipped\n"
                     % (len(manifest), covered, len(skipped)))
    json.dump({"resources": manifest, "skipped": skipped}, sys.stdout,
              indent=1, sort_keys=True)


if __name__ == "__main__":
    main()
