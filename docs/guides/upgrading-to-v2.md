---
subcategory: "Guides"
page_title: "Upgrading to v2"
description: |-
  What changed in v2.0.0 and the edits needed to move a v1 configuration across.
---

# Upgrading to v2

v2.0.0 corrects a set of attributes that addressed RouterOS properties which do
not exist, or that could not hold the values RouterOS actually reports. Every
change below affects configuration that **could not have worked on v1** — the
values were either rejected by the device or silently discarded — but several
need a mechanical edit before `terraform plan` will run.

Start by bumping the constraint:

```terraform
terraform {
  required_providers {
    routeros = {
      source  = "ebogdum/routeros"
      version = "~> 2.0"
    }
  }
}
```

Then work through the four categories below. Running `terraform plan` after each
one tells you whether anything is left.

## 1. Renamed attributes

The REST field names behind these were built by splitting an identifier, which
inserted hyphens RouterOS does not use — the provider sent `po-e-out` where the
device reports `poe-out`. Reads were wrong the same way, so these attributes
never populated *or* applied.

| v1 | v2 |
|---|---|
| `po_e_out` | `poe_out` |
| `po_e_priority` | `poe_priority` |
| `po_e_voltage` | `poe_voltage` |
| `po_e_out_current` | `poe_out_current` |
| `po_e_out_power` | `poe_out_power` |
| `po_e_out_status` | `poe_out_status` |
| `po_e_out_voltage` | `poe_out_voltage` |
| `don_t_require_permissions` | `dont_require_permissions` |

A find-and-replace across your configuration is sufficient. No state migration is
needed: the old attributes never held a device value.

## 2. Retyped attributes

RouterOS reports a word where the menu documents a number, so these could not
represent the router's own defaults. They are strings now:

| Resource | Attribute | Sentinel |
|---|---|---|
| `routeros_interface_bridge` | `mtu` | `auto` |
| `routeros_interface_bridge_port` | `horizon` | `none` |
| `routeros_ipv6_nd` | `hop_limit`, `mtu`, `reachable_time`, `retransmit_interval` | `unspecified` |

HCL converts a bare number to a string automatically, so `mtu = 1500` continues
to work unchanged. You only need an edit if you were computing these values in a
way that requires a number type — for example `mtu = var.mtu + 100`.

`routeros_interface_ethernet.auto_negotiation` changed the other way. In v1 it
was a string constrained to the read-only link state (`done`, `incomplete`,
`failed`), which is reported by `/interface/ethernet monitor` rather than the
menu, and it was never written to the device. In v2 it is the settable toggle:

```terraform
# v1 — inert, and rejected any value the device actually reported
auto_negotiation = "done"

# v2
auto_negotiation = true
```

## 3. Read-only attributes are no longer settable

561 attributes across 108 resources held read-only runtime state — `dynamic`,
`invalid`, `actual_mtu`, `fingerprint`, `owner` and similar — while advertising
themselves as `Optional`. Setting one was silently discarded.

They are `Computed`-only in v2. They still appear in state and remain readable as
`resource.x.dynamic`, but a value in configuration is now an error:

```
Error: Value for unconfigurable attribute
```

The fix is to delete the assignment. If you were setting one deliberately, note
that it never reached the device on v1 either.

## 4. Deprecated attributes

These are WebFig-only spellings that shadow a real property. RouterOS rejects
them over the REST API, so any configuration setting one failed to apply. They
still exist and still read, but are no longer sent, and will be removed in a
future release:

| Resource | Deprecated | Use instead |
|---|---|---|
| `routeros_interface_ethernet` | `autoneg`, `noautoneg` | `auto_negotiation` |
| `routeros_ip_hotspot_user` | `def`, `nondef` | `default` (read-only) |
| `routeros_ip_ipsec_mode_config` | `resp`, `nonresp` | `responder` |
| `routeros_ip_ipsec_policy` | `notemplate` | `template` |
| `routeros_ppp_profile` | `def` | `default` (read-only) |

## Wifi configuration now reaches the device

No edit is required for this, but it changes behaviour and is worth knowing
about before you apply.

RouterOS exposes the wifi sub-objects under dotted REST keys
(`configuration.ssid`, `security.passphrase`, `channel.band`). v1 read and wrote
them as flat names, which the device does not have, so **none of the wifi
settings round-tripped** — a CAP rebuilt from Terraform came up with no SSID and
no passphrase.

v2 maps each attribute to its dotted wire name. The Terraform attribute names are
unchanged, so your configuration does not need editing; it simply takes effect
now. Review your wifi plan before applying: values that were previously ignored
will be written to the device for the first time.

## New resources

v2 adds resources for 27 menus that previously had none, including
`routeros_ip_service` (enable/disable `telnet`, `ftp`, `api`, …),
`routeros_ip_dns_adlist`, and the RouterBOARD button bindings. See the resource
index for the full list. Nothing needs to change to keep using v1 resources.
