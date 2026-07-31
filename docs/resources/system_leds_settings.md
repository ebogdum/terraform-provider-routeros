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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `all_leds_off` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_leds_settings.this 'home'
```
