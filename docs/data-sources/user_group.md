---
subcategory: "Users"
page_title: "RouterOS: routeros_user_group"
description: |-
  RouterOS resource.
---

# Data Source: routeros_user_group

Manages the RouterOS `/user/group` menu.

## Example Usage

```terraform
data "routeros_user_group" "group_example" {
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
* `name` - (Required) Type: `string`. Default: `tf-acc-group`.
* `policy` - (Optional) Type: `list`. Default: `read`.
* `skin` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

