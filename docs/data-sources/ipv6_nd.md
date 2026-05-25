---
page_title: "RouterOS: routeros_ipv6_nd"
description: |-
  ND config is per-interface and conflicts with defaults if the interface is already configured.
---

# Data Source: routeros_ipv6_nd

ND config is per-interface and conflicts with defaults if the interface is already configured.

## Example Usage

```terraform
data "routeros_ipv6_nd" "nd_example" {
  # router   = "my-router"  # omit for the default router
  # filter   = { name = "some-name" }
  # proplist = ["name", "address"]
}
```

## Argument Reference

This data source supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to query.
* `filter` - (Optional) Map of field=value pairs to narrow the result set.
* `proplist` - (Optional) List of property names to project; smaller payload.
* `advertise_dns` - (Optional) Type: `enum(no|yes|self)`.
* `advertise_mac_address` - (Optional) Type: `bool`. Default: `1`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `hop_limit` - (Optional) Type: `int`. Default: `64`.
* `interface` - (Optional) Type: `string`.
* `managed_address_configuration` - (Optional) Type: `bool`.
* `mtu` - (Optional) Type: `int`.
* `other_configuration` - (Optional) Type: `bool`.
* `ra_delay` - (Optional) Type: `duration`. Default: `3`.
* `ra_interval` - (Optional) Type: `string`.
* `ra_lifetime` - (Optional) Type: `duration`. Default: `1800`.
* `ra_preference` - (Optional) Type: `enum(medium|high|low)`.
* `reachable_time` - (Optional) Type: `int`.
* `retransmit_interval` - (Optional) Type: `int`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `default` - Type: `bool`.
* `invalid` - Type: `bool`.

