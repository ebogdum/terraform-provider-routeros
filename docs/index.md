---
page_title: "Provider: RouterOS"
description: |-
  Manage MikroTik RouterOS 7.x devices through Terraform: every menu, every property, every action.
---

# RouterOS Provider

Manages MikroTik RouterOS 7.x devices over the device's REST API. It covers
the menus RouterOS exposes: 310 resources, 285 data sources and 75 actions,
with each property's schema validated against a live router.

One provider block can manage a single router or a whole fleet through a
named `routers` map. Every resource, data source and action takes an optional
`router` attribute, which you leave out to target the default.

## Before You Start

RouterOS is not one API. It is one per board and one per release. A property
that exists on a hAP ax³ running 7.23 can be missing on a CCR, read-only on a
CRS, or renamed in the next stable. Whole menus for hardware you do not have
(CAPsMAN, wireless, LTE, UPS, SFP) are simply absent. The schema here was
validated against real hardware, but no single device exposes all of it, and
MikroTik keeps moving things between releases.

So expect the occasional rough edge on a board or a RouterOS version that
differs from the one a given property was checked against. It usually shows up
as a 400 from the device on apply, an attribute that never stops showing a
diff, or a value that reads back empty.

If you run into one, please open a pull request. The fix is usually a few lines
in that resource's schema, and a change that comes with the RouterOS version,
the board name and the exact error the device returned is most of the way
there already.

## Example Usage

```terraform
terraform {
  required_providers {
    routeros = {
      source  = "ebogdum/routeros"
      version = "~> 3.0"
    }
  }
}

provider "routeros" {
  host     = "https://192.0.2.1"
  username = "admin"
  password = var.routeros_password
  insecure = true
}

resource "routeros_ip_address" "lan" {
  address   = "192.168.88.1/24"
  interface = "bridge1"
  comment   = "LAN gateway"
}
```

## Multi-Router Fleet

```terraform
provider "routeros" {
  routers = {
    core = {
      host     = "https://10.0.0.1"
      username = "admin"
      password = var.core_password
      insecure = true
    }
    edge_se = {
      host     = "https://10.0.1.1"
      username = "admin"
      password = var.edge_se_password
      insecure = true
    }
  }
}

resource "routeros_system_identity" "label" {
  for_each = toset(["core", "edge_se"])
  router   = each.key
  name     = each.key
}

data "routeros_ip_route" "core_routes" {
  router = "core"
}
```

## Environment Variables

Single-router shorthand fields fall back to environment variables when not
set in the provider block:

| Variable | Provider attribute |
|---|---|
| `ROUTEROS_HOST` | `host` |
| `ROUTEROS_USER` | `username` |
| `ROUTEROS_PASSWORD` | `password` |
| `ROUTEROS_CA_CERT` | `ca_cert` |
| `ROUTEROS_INSECURE` | `insecure` |
| `ROUTEROS_VERSION` | `ros_version` |

## Safety Guards

Changes that would obviously sever management access are refused unless
the resource sets `lockout_ack = true`:

- `routeros_ip_firewall_filter` / `routeros_ipv6_firewall_filter` --
  `chain=input|forward` + `action=drop|reject|tarpit` with no narrowing match
- `routeros_user` -- deleting or disabling the last admin
- `routeros_user_group` -- removing required policies from the `full` group
- `routeros_tool_mac_server` -- emptying `allowed-interface-list`

Sensitive properties (passwords, secrets, keys, OTP, SIM PIN) are flagged
so Terraform redacts them in plan output, state files, and CLI display.
