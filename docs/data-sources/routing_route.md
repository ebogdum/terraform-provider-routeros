---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_route"
description: |-
  RouterOS resource.
---

# Data Source: routeros_routing_route

Manages the RouterOS `/routing/route` menu.

## Example Usage

```terraform
data "routeros_routing_route" "route_example" {
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

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `active` - Type: `bool`.
* `afi` - Type: `string`.
* `belongs_to` - Type: `string`.
* `connect` - Type: `bool`.
* `contribution` - Type: `string`.
* `debug_fwp_ptr` - Type: `int`.
* `dhcp` - Type: `bool`.
* `distance` - Type: `int`.
* `dst_address` - Type: `cidr`.
* `gateway` - Type: `ip`.
* `immediate_gw` - Type: `string`.
* `local_address` - Type: `string`.
* `nexthop_id` - Type: `string`.
* `routing_table` - Type: `string`.
* `scope` - Type: `int`.
* `target_scope` - Type: `int`.
* `vrf_interface` - Type: `string`.

