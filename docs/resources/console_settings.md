---
subcategory: "System & misc"
page_title: "RouterOS: routeros_console_settings"
description: |-
  RouterOS resource.
---

# Resource: routeros_console_settings

Manages the RouterOS `/console/settings` menu.

## Example Usage

```terraform
resource "routeros_console_settings" "settings_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # log_script_errors = false
  # sanitize_names = false
  # tab_width = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `log_script_errors` - (Optional) Type: `bool`.
* `sanitize_names` - (Optional) Type: `bool`.
* `tab_width` - (Optional) Type: `int`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_console_settings.this 'home'
```
