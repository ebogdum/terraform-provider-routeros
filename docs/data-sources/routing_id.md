---
page_title: "RouterOS: routeros_routing_id"
description: |-
  RouterOS resource.
---

# Data Source: routeros_routing_id

Manages the RouterOS `/routing/id` menu.

## Example Usage

```terraform
data "routeros_routing_id" "id_example" {
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
* `name` - (Optional) Type: `string`.
* `select_dynamic_id` - (Optional) Type: `enum(only static|only loopback|only vrf|only active|any|lowest)`.
* `select_from_vrf` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `dynamic` - Type: `bool`.
* `dynamic_id` - Type: `ip`.
* `inactive` - Type: `bool`.

