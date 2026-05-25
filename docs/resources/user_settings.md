---
page_title: "RouterOS: routeros_user_settings"
description: |-
  RouterOS resource.
---

# Resource: routeros_user_settings

Manages the RouterOS `/user/settings` menu.

## Example Usage

```terraform
resource "routeros_user_settings" "settings_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # minimum_categories = 0
  # minimum_password_length = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `minimum_categories` - (Optional) Type: `int`.
* `minimum_password_length` - (Optional) Type: `int`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_user_settings.this 'home'
```
