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
| Whether a menu holds `none` / `one` / `many` rows | How many, or which |
| Board model, RouterOS version, architecture, firmware | Serial number, licence, system identity |

Row presence is bucketed to `none`, `one` or `many` on purpose. Whether a menu
holds one row or several is what separates a settings singleton from a
per-port collection — a distinction the provider gets wrong on
`/interface/ethernet/poe` — but the exact number is nobody's business: how many
firewall connections or bridge hosts you have describes your network. Pass
`--no-rows` to skip reading rows altogether.

After a full run against a hAP ax³ the only digit anywhere in the output is the
collector version.

`--with-enums` additionally records the value sets RouterOS offers for
enum-shaped arguments, such as `auto` / `high` / `low`. Those come from the
console's completion list, not from your configuration.

Before writing anything out, the script checks its own output: every recorded
name must match `[A-Za-z0-9._-]+`, so a value cannot ride along inside one.

**It never writes to the device.** Every command is an inspect or a print.

## Running it

Python 3, no dependencies. It talks REST by default — the same API the provider
uses — so a username and password is all you need.

```sh
ROUTEROS_HOST=192.168.88.1 ROUTEROS_USER=admin ROUTEROS_PASSWORD=... \
  python3 tools/conformance/collect_schema.py -o my-board.json
```

or, from the repository root:

```sh
ROUTEROS_HOST=192.168.88.1 ROUTEROS_PASSWORD=... make collect COLLECT_ARGS="-o my-board.json"
```

A full tree is about 30 seconds and 320 KB. If the `www` service is switched
off, `--transport ssh` uses the console over an SSH key instead; it collects the
same menus and arguments but fewer row property names, and takes several
minutes.

Useful flags:

```sh
--check          # print a summary instead of JSON, to see what it would send
--no-rows        # skip reading rows entirely
--root /interface/ethernet   # one subtree only
--transport ssh  # console over SSH instead of REST
--insecure       # accept a self-signed certificate on https
--jobs N         # concurrent requests, default 4
-o FILE          # write here instead of stdout
```

Look before you send:

```sh
ROUTEROS_HOST=192.168.88.1 python3 tools/conformance/collect_schema.py --check
```

```
RB5009UPr+S+  7.23.3 (stable)  arm64  firmware 7.19.6  (via rest)
437 menus, 4228 settable arguments, 1685 row properties
171 menus hold rows
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
