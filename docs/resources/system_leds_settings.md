---
subcategory: "System"
page_title: "RouterOS: routeros_system_leds_settings"
description: |-
  RouterOS resource.
---

# Resource: routeros_system_leds_settings

Manages the RouterOS `/system/leds/settings` menu.

## Example Usage

```terraform
resource "routeros_system_leds_settings" "settings_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # all_leds_off = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `all_leds_off` - (Optional) Type: `string`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_leds_settings.this 'home'
```
