---
subcategory: "System"
page_title: "RouterOS: routeros_system_device_mode_update_mode"
description: |-
  RouterOS resource.
---

# Data Source: routeros_system_device_mode_update_mode

Manages the RouterOS `/system/device-mode/update/mode` menu.

## Example Usage

```terraform
data "routeros_system_device_mode_update_mode" "mode_example" {
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
* `get` - (Optional) Type: `string`. Returns value that you can assign to variable or print on the screen.
* `print` - (Optional) Type: `string`. Shows the active mode and its properties.
* `update` - (Optional) Type: `string`. Applies changes to the specified properties, see below.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

