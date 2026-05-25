---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_rip_neighbor"
description: |-
  Discovered; needs rip neighbor fixture
---

# Data Source: routeros_routing_rip_neighbor

Discovered; needs rip neighbor fixture

## Example Usage

```terraform
data "routeros_routing_rip_neighbor" "neighbor_example" {
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
* `disabled` - (Optional) Type: `bool`.
* `instance` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

