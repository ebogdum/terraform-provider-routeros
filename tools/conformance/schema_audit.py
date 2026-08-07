#!/usr/bin/env python3
"""Diff the provider's RouterOS property names against a live device's schema.

Tests only check what they exercise. This checks the whole surface: every
property name the provider can put on the wire, compared against the argument
list RouterOS itself reports. That is how the v3.0.4/v3.0.5 defects were found
after a green test suite -- an acceptance test that sets no attributes never
puts a wrong name on the wire, and a wrong read key is a silently skipped
branch that leaves the attribute null (legal for Optional+Computed).

Two independent checks, both sourced from the device:

  names   Every body["x"] / obj["x"] literal in the provider, against
          `/console/inspect request=child path=<menu>,set`, the authoritative
          list of arguments RouterOS accepts. A written name that is not in
          that list is a guaranteed HTTP 400 for anyone who sets it.

  types   Every attribute the provider declares as bool, against
          `/console/inspect request=completion path=<menu>,set,<prop>`.
          A genuine bool completes to exactly yes/no. Anything else is an
          enum (or a reference) that a bool schema cannot express.

Menus absent from the running hardware are reported separately rather than
counted as failures -- no single board exposes every menu.

Usage:
    ROUTEROS_HOST=192.168.10.2 ROUTEROS_USER=admin ROUTEROS_PASSWORD=... \
        python3 tools/conformance/schema_audit.py [names|types|all]

Exit status is 1 when a check finds a defect, so it can gate a release.
"""

import os
import re
import subprocess
import sys
from glob import glob

HOST = os.environ.get("ROUTEROS_HOST", "").replace("http://", "").replace("https://", "")
USER = os.environ.get("ROUTEROS_USER", "admin")
PASSWORD = os.environ.get("ROUTEROS_PASSWORD", "")
SRC = "internal/provider"
BATCH = 25


def ssh(script, timeout=240):
    """Run a RouterOS console script over SSH and return stdout."""
    return subprocess.run(
        ["ssh", "-o", "BatchMode=yes", "-o", "ConnectTimeout=15",
         "%s@%s" % (USER, HOST), script],
        capture_output=True, text=True, timeout=timeout).stdout


def console_path(menu, *rest):
    return ",".join([menu.strip("/").replace("/", ",")] + list(rest))


def provider_properties():
    """Every (menu, prop) the provider reads or writes, from the source."""
    out, ambiguous = [], []
    files = sorted(glob(SRC + "/resource_*.go") + glob(SRC + "/action_*.go"))
    for f in files:
        if f.endswith("_test.go"):
            continue
        src = open(f).read()
        menus = sorted(set(re.findall(r'ctx,\s*"(/[a-z0-9/\-]+)"', src)))
        written = set(re.findall(r'body\["([^"]+)"\]\s*=', src))
        read = set(re.findall(r'obj\["([^"]+)"\]', src))
        props = (written | read) - {".id", ".nextid"}
        if not props:
            continue
        if len(menus) != 1:
            ambiguous.append(os.path.basename(f))
            continue
        for p in sorted(props):
            out.append({"file": os.path.basename(f), "menu": menus[0], "prop": p,
                        "written": p in written})
    return out, ambiguous


def device_set_args(menus):
    """Authoritative `set` argument names per menu, straight off the device."""
    args = {}
    for i in range(0, len(menus), BATCH):
        cmds = []
        for m in menus[i:i + BATCH]:
            cmds.append(':put "@@@%s"' % m)
            cmds.append(':put [:tostr [/console/inspect request=child '
                        'path=%s as-value]]' % console_path(m, "set"))
        cur = None
        for line in ssh("; ".join(cmds)).splitlines():
            line = line.strip()
            if line.startswith("@@@"):
                cur = line[3:]
                args.setdefault(cur, [])
            elif cur and "name=" in line:
                args[cur] = re.findall(r"name=([^;]+);node-type=arg", line)
        sys.stderr.write("  args %d/%d\n" % (i // BATCH + 1,
                                             (len(menus) + BATCH - 1) // BATCH))
    return args


def device_completions(pairs):
    """Accepted values for each (menu, prop), straight off the device."""
    res = {}
    for i in range(0, len(pairs), BATCH):
        cmds = []
        for menu, prop in pairs[i:i + BATCH]:
            cmds.append(':put "@@@%s|%s"' % (menu, prop))
            cmds.append(':put [:tostr [/console/inspect request=completion '
                        'path=%s as-value]]' % console_path(menu, "set", prop))
        cur = None
        for line in ssh("; ".join(cmds)).splitlines():
            line = line.strip()
            if line.startswith("@@@"):
                cur = tuple(line[3:].split("|", 1))
                res.setdefault(cur, [])
            elif cur and "completion=" in line:
                res[cur] = [v for v, _, show, _ in re.findall(
                    r"completion=(.*?);offset=\d+;preference=(-?\d+);"
                    r"show=(\w+);style=(\S*?);", line) if show == "true"]
        sys.stderr.write("  values %d/%d\n" % (i // BATCH + 1,
                                               (len(pairs) + BATCH - 1) // BATCH))
    return res


def bool_attributes():
    """(menu, prop) pairs the provider declares as a Terraform bool."""
    out = []
    for f in sorted(glob(SRC + "/resource_*.go")):
        if f.endswith("_test.go"):
            continue
        src = open(f).read()
        if "schema.BoolAttribute" not in src:
            continue
        menus = sorted(set(re.findall(r'ctx,\s*"(/[a-z0-9/\-]+)"', src)))
        if len(menus) != 1:
            continue
        props = set(re.findall(
            r'body\["([^"]+)"\]\s*=\s*client\.FormatBool\(', src))
        for m in re.finditer(r'obj\["([^"]+)"\]', src):
            chunk = src[m.end(): m.end() + 400]
            nxt = chunk.find('obj["')
            if nxt != -1:
                chunk = chunk[:nxt]
            if "types.BoolValue" in chunk or "types.BoolNull" in chunk:
                props.add(m.group(1))
        for p in sorted(props):
            out.append((menus[0], p))
    return sorted(set(out))


def check_names():
    props, ambiguous = provider_properties()
    menus = sorted({p["menu"] for p in props})
    sys.stderr.write("names: %d properties across %d menus\n"
                     % (len(props), len(menus)))
    args = device_set_args(menus)

    broken, unreachable = [], set()
    for p in props:
        a = args.get(p["menu"]) or []
        if not a:
            unreachable.add(p["menu"])
            continue
        if p["written"] and p["prop"] not in a:
            broken.append(p)

    print("\n== write-name check ==")
    print("menus not present on this hardware (skipped): %d" % len(unreachable))
    if ambiguous:
        print("files whose menu path could not be resolved: %s"
              % ", ".join(ambiguous))
    print("WRITTEN names the device rejects: %d" % len(broken))
    for p in sorted(broken, key=lambda x: (x["menu"], x["prop"])):
        near = [x for x in args[p["menu"]] if x.endswith("." + p["prop"])]
        hint = ("  -> %s" % near[0]) if len(near) == 1 else ""
        print("  %-46s %-34s (%s)%s" % (p["menu"], p["prop"], p["file"], hint))
    return len(broken)


def check_types():
    pairs = bool_attributes()
    sys.stderr.write("types: %d bool attributes\n" % len(pairs))
    comp = device_completions(pairs)

    mismatched, unverifiable = [], 0
    for (menu, prop), vals in sorted(comp.items()):
        if not vals:
            unverifiable += 1
            continue
        if set(vals) != {"yes", "no"}:
            mismatched.append((menu, prop, vals))

    print("\n== bool-vs-enum check ==")
    print("bool attributes not settable here (read-only flags / absent "
          "menus): %d" % unverifiable)
    print("declared bool but the device accepts other values: %d"
          % len(mismatched))
    for menu, prop, vals in mismatched:
        print("  %-46s %-30s %s" % (menu, prop, vals))
    return len(mismatched)


def main():
    if not HOST:
        sys.exit("ROUTEROS_HOST is required (e.g. 192.168.10.2)")
    which = sys.argv[1] if len(sys.argv) > 1 else "all"
    failures = 0
    if which in ("names", "all"):
        failures += check_names()
    if which in ("types", "all"):
        failures += check_types()
    print("\n%s" % ("AUDIT CLEAN" if failures == 0
                    else "AUDIT FOUND %d DEFECT(S)" % failures))
    sys.exit(1 if failures else 0)


if __name__ == "__main__":
    main()
