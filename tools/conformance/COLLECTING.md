# Collecting a device schema

`collect_schema.py` walks a RouterOS device's menu tree and writes down its
**shape** — which menus exist, what arguments each one accepts, which properties
appear in its rows. It never records a value.

If you run MikroTik hardware this provider has not been tested against, running
this and attaching the JSON to an issue is the single most useful thing you can
contribute.

## Why it is needed

The provider's schema was generated against one RouterOS version on one board.
Every board disagrees with it somewhere. A hAP ax³ has no disk, no PoE-in, no SFP
cage and a limited switch chip, so large parts of `/disk`,
`/interface/ethernet/poe` and `/interface/ethernet/switch` cannot be checked on
it at all — which is exactly where the provider's wrong property names have been
hiding. Some names differ *between* boards: `mirror-source` is
`mirror-egress-target` on an RB5009 and `mirror-target` on a CRS305.

There is no way to know any of that without reading the schema off hardware that
has the parts.

## What it sends

**Names and structure only.**

| Recorded | Not recorded |
| --- | --- |
| Menu paths (`/ip/firewall/filter`) | Any property value |
| Whether a menu is a namespace or a real menu | Addresses, passwords, keys, certificates |
| Argument names accepted by `set` and `add` | Interface names, SSIDs, comments |
| Command names (`print`, `move`, `reset`) | Firewall rules, routes, DNS entries |
| Property names appearing in rows | Row contents |
| How many rows a menu has | Which rows, or anything in them |
| Board model, RouterOS version, architecture, firmware | Serial number, licence, identity |

Row *counts* are included because "this menu has 8 rows" is what distinguishes a
per-port collection from a settings singleton — a distinction the provider
currently gets wrong on `/interface/ethernet/poe`. If even that is more than you
want to share, pass `--no-rows`.

`--with-enums` additionally records the value sets RouterOS offers for
enum-shaped arguments, such as `auto` / `high` / `low`. Those come from the
console's completion list, not from your configuration.

Before writing anything out, the script checks its own output: every recorded
name must match `[A-Za-z0-9._-]+`, so a value cannot ride along inside one.

**It never writes to the device.** Every command is an inspect or a print.

## Running it

Needs Python 3 and SSH key access. No password is read or stored.

```sh
ROUTEROS_HOST=192.168.88.1 ROUTEROS_USER=admin \
  python3 tools/conformance/collect_schema.py > my-board.json
```

or, from the repository root:

```sh
ROUTEROS_HOST=192.168.88.1 make collect > my-board.json
```

A full tree is around 30 seconds and 300 KB.

Useful flags:

```sh
--check        # print a summary instead of JSON, to see what it would send
--no-rows      # skip reading rows entirely
--with-enums   # also record enum value sets (slower)
--root /interface/ethernet   # one subtree only
```

Look before you send:

```sh
ROUTEROS_HOST=192.168.88.1 python3 tools/conformance/collect_schema.py --check
```

```
RB5009UPr+S+  7.23.3 (stable)  arm64  firmware 7.19.6
437 menus, 4228 settable arguments, 1370 row properties
155 menus have at least one row
```

## Sending it

Open an issue titled with the board and version — for example
`RB5009UPr+S+ 7.23.3 schema` — and attach the JSON. Boards with hardware the
common test units lack are the most valuable: PoE-in or PoE-out, an SFP or SFP+
cage, a real switch chip, a disk or USB storage, LTE, or 60 GHz.

Please say which RouterOS channel you are on (stable, testing, long-term), since
menus move between releases.

## What happens to it

The JSON is compared against the provider's declared schema to find:

- properties the provider writes that the board does not accept
- properties the board accepts that the provider does not expose
- names that differ between boards, which need per-board handling
- menus that exist only on some hardware
- submenus with no resource behind them at all

`schema_audit.py` does that comparison live against a device you own. The
collector exists so the same comparison can be made against hardware the
maintainers do not have.
