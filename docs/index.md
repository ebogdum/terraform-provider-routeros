---
page_title: "Provider: routeros"
description: |-
  Manage MikroTik RouterOS 7.x devices through Terraform.
---

# routeros Provider

Manage MikroTik RouterOS 7.x devices through Terraform.

The provider speaks RouterOS REST (HTTPS) and covers every menu the device
exposes: IP, IPv6, routing, firewall, interfaces, queues, certificates,
hotspot, PPP, IPsec, system, tools, and more -- 420 menus / 182 resources /
276 data sources / 75 actions, all generated from a single committed schema.

One provider block can manage many routers at once via the `routers` map;
each resource and data source picks the target router with an optional
`router` attribute.

## Example Usage -- single router

```terraform
terraform {
  required_providers {
    routeros = {
      source  = "ebogdum/routeros"
      version = "~> 1.0"
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

## Example Usage -- multi-router (manage a fleet from one config)

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
    edge_nw = {
      host     = "https://10.0.2.1"
      username = "admin"
      password = var.edge_nw_password
      insecure = true
    }
  }
}

# Apply a uniform identity to every router with a single resource block.
resource "routeros_system_identity" "label" {
  for_each = toset(["core", "edge_se", "edge_nw"])
  router   = each.key
  name     = each.key
}

# Cross-router lookup: read core's routing table from edge_se's config.
data "routeros_ip_route" "core_routes" {
  router = "core"
}
```

Omitting `router =` selects the default -- the entry named `default`, or the
first router in sorted order if there is no `default`.

## Schema

### Single-router shorthand

| field | type | required | description |
|---|---|---|---|
| `host` | string | no (env: `ROUTEROS_HOST`) | Router base URL, e.g. `https://192.0.2.1`. |
| `username` | string | no (env: `ROUTEROS_USER`) | API user. |
| `password` | string | no (env: `ROUTEROS_PASSWORD`) | API password. **Sensitive.** |
| `ca_cert` | string | no (env: `ROUTEROS_CA_CERT`) | PEM-encoded CA bundle for verifying the router's TLS cert. |
| `insecure` | bool | no (env: `ROUTEROS_INSECURE`) | Skip TLS verification. Use only for lab gear. |
| `ros_version` | string | no (env: `ROUTEROS_VERSION`) | Declared RouterOS version; used to gate version-conditional fields. |

Shorthand fields seed a router named `default`. They are ignored when
`routers = { default = { ... } }` is already declared.

### Multi-router

```hcl
provider "routeros" {
  routers = {
    <name> = {
      host     = "..."
      username = "..."
      password = "..."     # sensitive
      ca_cert  = "..."     # optional
      insecure = true|false
    }
    ...
  }
}
```

## Safety

The provider refuses to apply changes that would obviously sever management
access to the router, unless the resource sets `lockout_ack = true`:

- `routeros_ip_firewall_filter` / `routeros_ipv6_firewall_filter` --
  `chain=input|forward` with `action=drop|reject|tarpit` and no narrowing
  match condition (the rule would drop your own session).
- `routeros_user` -- deleting or disabling the last admin account.
- `routeros_user_group` -- removing required policies (`api`, `web`, `winbox`,
  `ssh`, `password`, `policy`, `write`) from the `full` group.
- `routeros_tool_mac_server` -- emptying the `allowed-interface-list`.

Each guard is conservative: the override is per-resource, not provider-wide.

## Status

See the [CHANGELOG](https://github.com/ebogdum/terraform-provider-routeros/blob/main/CHANGELOG.md)
for release history. Schema is harvested + validated against a live device
per release; menus that require hardware not present on the test device
are emitted but not exercised in the automated acceptance sweep.
