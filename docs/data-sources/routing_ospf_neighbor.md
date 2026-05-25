---
subcategory: "OSPF"
page_title: "RouterOS: routeros_routing_ospf_neighbor"
description: |-
  RouterOS resource.
---

# Data Source: routeros_routing_ospf_neighbor

Manages the RouterOS `/routing/ospf/neighbor` menu.

## Example Usage

```terraform
data "routeros_routing_ospf_neighbor" "neighbor_example" {
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
* `address` - (Optional) Type: `string`.
* `area` - (Optional) Type: `string`.
* `bdr` - (Optional) Type: `ip`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `dr` - (Optional) Type: `ip`.
* `instance` - (Optional) Type: `string`.
* `priority` - (Optional) Type: `int`.
* `router_id` - (Optional) Type: `ip`.
* `state` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `adjacency` - Type: `string`.
* `db_summaries` - Type: `int`.
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.
* `ls_requests` - Type: `int`.
* `ls_retransmits` - Type: `int`.
* `state_changes` - Type: `int`.

