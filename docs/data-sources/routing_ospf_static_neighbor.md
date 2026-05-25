---
subcategory: "OSPF"
page_title: "RouterOS: routeros_routing_ospf_static_neighbor"
description: |-
  References an existing ospf area; auto-test can't synthesise.
---

# Data Source: routeros_routing_ospf_static_neighbor

References an existing ospf area; auto-test can't synthesise.

## Example Usage

```terraform
data "routeros_routing_ospf_static_neighbor" "static_neighbor_example" {
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
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `instance_id` - (Optional) Type: `int`.
* `poll_interval` - (Optional) Type: `duration`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

