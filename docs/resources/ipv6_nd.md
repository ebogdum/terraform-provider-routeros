---
subcategory: "IPv6"
page_title: "RouterOS: routeros_ipv6_nd"
description: |-
  ND config is per-interface and conflicts with defaults if the interface is already configured.
---

# Resource: routeros_ipv6_nd

ND config is per-interface and conflicts with defaults if the interface is already configured.

## Example Usage

```terraform
resource "routeros_ipv6_nd" "nd_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # advertise_dns = "no"
  # advertise_mac_address = "10.99.0.0/24"
  # dns_servers = "replace-me"
  # hop_limit = 64
  # interface = "ether1"
  # managed_address_configuration = false
  # mtu = 0
  # other_configuration = false
  # pref64_prefixes = "replace-me"
  # ra_delay = "3"
  # ra_interval = "replace-me"
  # ra_lifetime = "1800"
  # ra_preference = "medium"
  # reachable_time = 0
  # retransmit_interval = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `advertise_dns` - (Optional) Type: `string`.
* `advertise_mac_address` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default` - (Read-only) Type: `bool`.
* `disabled` - (Optional) Type: `bool`.
* `dns` - (Optional) Type: `string`. RouterOS `dns`.
* `dns_servers` - (Read-only) Type: `string`.
* `hop_limit` - (Optional) Type: `string`. Hop limit advertised in router advertisements. A number, or `unspecified` (the default).
* `interface` - (Optional) Type: `string`.
* `invalid` - (Read-only) Type: `bool`.
* `managed_address_configuration` - (Optional) Type: `bool`.
* `mtu` - (Optional) Type: `string`. MTU advertised in router advertisements. A number, or `unspecified` (the default).
* `other_configuration` - (Optional) Type: `bool`.
* `pref64` - (Optional) Type: `string`. RouterOS `pref64`.
* `pref64_prefixes` - (Read-only) Type: `string`.
* `ra_delay` - (Optional) Type: `string`.
* `ra_interval` - (Optional) Type: `string`.
* `ra_lifetime` - (Optional) Type: `string`.
* `ra_preference` - (Optional) Type: `string`.
* `reachable_time` - (Optional) Type: `string`. Reachable time advertised in router advertisements. A number, or `unspecified` (the default).
* `retransmit_interval` - (Optional) Type: `string`. Retransmit interval advertised in router advertisements. A number, or `unspecified` (the default).

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ipv6_nd.example '*3'

# Named router
terraform import routeros_ipv6_nd.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ipv6_nd.example 'home/my-resource-name'
```
