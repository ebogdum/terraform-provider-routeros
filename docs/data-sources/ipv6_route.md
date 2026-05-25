---
subcategory: "IPv6"
page_title: "RouterOS: routeros_ipv6_route"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ipv6_route

Manages the RouterOS `/ipv6/route` menu.

## Example Usage

```terraform
data "routeros_ipv6_route" "route_example" {
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
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `distance` - (Optional) Type: `int`.
* `dst_address` - (Optional) Type: `cidr`.
* `gateway` - (Optional) Type: `string`.
* `routing_table` - (Optional) Type: `string`.
* `scope` - (Optional) Type: `int`.
* `target_scope` - (Optional) Type: `int`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `active` - Type: `bool`.
* `connect` - Type: `bool`.
* `dynamic` - Type: `bool`.
* `immediate_gw` - Type: `string`.
* `inactive` - Type: `bool`.

