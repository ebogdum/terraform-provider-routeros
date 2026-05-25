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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `advertise_dns` - (Optional) Type: `enum(no|yes|self)`.
* `advertise_mac_address` - (Optional) Type: `bool`. Default: `1`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `dns_servers` - (Optional) Type: `string`.
* `hop_limit` - (Optional) Type: `int`. Default: `64`.
* `interface` - (Optional) Type: `string`.
* `managed_address_configuration` - (Optional) Type: `bool`.
* `mtu` - (Optional) Type: `int`.
* `other_configuration` - (Optional) Type: `bool`.
* `pref64_prefixes` - (Optional) Type: `string`.
* `ra_delay` - (Optional) Type: `duration`. Default: `3`.
* `ra_interval` - (Optional) Type: `string`.
* `ra_lifetime` - (Optional) Type: `duration`. Default: `1800`.
* `ra_preference` - (Optional) Type: `enum(medium|high|low)`.
* `reachable_time` - (Optional) Type: `int`.
* `retransmit_interval` - (Optional) Type: `int`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `default` - Type: `bool`.
* `invalid` - Type: `bool`.

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
