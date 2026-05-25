---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_filter_chain"
description: |-
  Discovered; required chain name must be unique
---

# Data Source: routeros_routing_filter_chain

Discovered; required chain name must be unique

## Example Usage

```terraform
data "routeros_routing_filter_chain" "chain_example" {
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
* `name` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.

