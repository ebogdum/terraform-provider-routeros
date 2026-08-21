---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_user_manager_router"
description: |-
  RouterOS resource.
---

# Data Source: routeros_user_manager_router

Manages the RouterOS `/user-manager/router` menu.

## Example Usage

```terraform
data "routeros_user_manager_router" "router_example" {
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
* `coa_port` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `protocol` - (Optional) Type: `string`.
* `shared_secret` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`. **Marked sensitive**: this menu holds a secret, which RouterOS returns in the row like any other column, so an unprojected read puts it in your state file. Use `proplist` to name the columns you need.

