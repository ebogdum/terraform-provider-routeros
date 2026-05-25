---
page_title: "RouterOS: routeros_interface_list"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_list

Manages the RouterOS `/interface/list` menu.

## Example Usage

```terraform
data "routeros_interface_list" "list_example" {
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
* `comment` - (Optional) Type: `string`.
* `exclude` - (Optional) Type: `string`. Defines interface list which members are excluded from the list. It is possible to add multiple lists separated by commas.
* `include` - (Optional) Type: `string`. Defines interface list which members are included in the list. It is possible to add multiple lists separated by commas.
* `name` - (Required) Type: `string`. Name of the interface list. Default: `tf_acc_iflist`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `builtin` - Type: `bool`.
* `dynamic` - Type: `bool`.

