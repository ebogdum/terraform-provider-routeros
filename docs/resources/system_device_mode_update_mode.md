---
subcategory: "System"
page_title: "RouterOS: routeros_system_device_mode_update_mode"
description: |-
  RouterOS resource.
---

# Resource: routeros_system_device_mode_update_mode

Manages the RouterOS `/system/device-mode/update/mode` menu.

## Example Usage

```terraform
resource "routeros_system_device_mode_update_mode" "mode_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # get = "replace-me"
  # print = "replace-me"
  # update = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `get` - (Optional) Type: `string`. Returns value that you can assign to variable or print on the screen.
* `print` - (Optional) Type: `string`. Shows the active mode and its properties.
* `update` - (Optional) Type: `string`. Applies changes to the specified properties, see below.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_system_device_mode_update_mode.example '*3'

# Named router
terraform import routeros_system_device_mode_update_mode.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_system_device_mode_update_mode.example 'home/my-resource-name'
```
