---
subcategory: "System"
page_title: "RouterOS: routeros_system_package"
description: |-
  Installed RouterOS packages. Managed by /system/upgrade and /system/package/update actions, not directly.
---

# Data Source: routeros_system_package

Installed RouterOS packages. Managed by /system/upgrade and /system/package/update actions, not directly.

## Example Usage

```terraform
data "routeros_system_package" "package_example" {
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
* `apply_changes` - (Optional) Type: `string`.
* `available` - (Optional) Type: `bool`.
* `build_time` - (Optional) Type: `string`.
* `bundle` - (Optional) Type: `int`.
* `check_for_updates` - (Optional) Type: `string`.
* `check_installation` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disable` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `downgrade` - (Optional) Type: `string`.
* `enable` - (Optional) Type: `string`.
* `installed` - (Optional) Type: `bool`.
* `name` - (Optional) Type: `string`.
* `scheduled` - (Optional) Type: `enum(|scheduled-for-uninstall|scheduled-for-disable|scheduled-for-enable|scheduled-for-install)`.
* `size` - (Optional) Type: `int`.
* `uninstall` - (Optional) Type: `string`.
* `unschedule` - (Optional) Type: `string`.
* `version` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

