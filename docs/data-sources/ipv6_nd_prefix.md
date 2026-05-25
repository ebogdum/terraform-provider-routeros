---
subcategory: "IPv6"
page_title: "RouterOS: routeros_ipv6_nd_prefix"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ipv6_nd_prefix

Manages the RouterOS `/ipv6/nd/prefix` menu.

## Example Usage

```terraform
data "routeros_ipv6_nd_prefix" "prefix_example" {
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
* `x6to4_interface` - (Optional) Type: `string`. Default: `4.294967295e+09`.
* `autonomous` - (Optional) Type: `bool`. Default: `1`.
* `dhcpv6_pd_preferred` - (Optional) Type: `bool`.
* `disabled` - (Optional) Type: `bool`.
* `interface` - (Optional) Type: `string`.
* `no6to4` - (Optional) Type: `string`.
* `on_link` - (Optional) Type: `bool`. Default: `1`.
* `preferred_lifetime` - (Optional) Type: `duration`. Default: `604800`.
* `prefix` - (Optional) Type: `string`.
* `valid_lifetime` - (Optional) Type: `duration`. Default: `2.592e+06`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.

