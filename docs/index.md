---
page_title: "Provider: RouterOS"
description: |-
  Manage MikroTik RouterOS 7.x devices through Terraform: every menu, every property, every action.
---

# RouterOS Provider

The RouterOS provider manages MikroTik RouterOS 7.x devices through their
REST API. It covers every menu the device exposes -- 186 resources, 277
data sources, 75 actions, 3522 properties -- all generated from a schema
validated property by property against a live router.

A single provider block can manage one router or an entire fleet via a
named `routers` map; every resource and data source takes an optional
`router` attribute, omitted for the default router.

## Example Usage

```terraform
terraform {
  required_providers {
    routeros = {
      source  = "ebogdum/routeros"
      version = "~> 2.0"
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
