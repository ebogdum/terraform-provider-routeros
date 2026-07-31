---
subcategory: "Users & RADIUS"
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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `minimum_categories` - (Optional) Type: `int`.
* `minimum_password_length` - (Optional) Type: `int`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_user_settings.this 'home'
```
