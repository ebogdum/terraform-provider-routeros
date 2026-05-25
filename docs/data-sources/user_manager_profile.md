---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_user_manager_profile"
description: |-
  RouterOS resource.
---

# Data Source: routeros_user_manager_profile

Manages the RouterOS `/user-manager/profile` menu.

## Example Usage

```terraform
data "routeros_user_manager_profile" "profile_example" {
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

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

