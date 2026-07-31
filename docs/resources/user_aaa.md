---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_user_aaa"
description: |-
  RouterOS resource.
---

# Resource: routeros_user_aaa

Manages the RouterOS `/user/aaa` menu.

## Example Usage

```terraform
resource "routeros_user_aaa" "aaa_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # accounting = false
  # default_group = "replace-me"
  # exclude_groups = "replace-me"
  # interim_update = "1h"
  # use_radius = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `accounting` - (Optional) Type: `bool`.
* `default_group` - (Optional) Type: `string`.
* `exclude_groups` - (Optional) Type: `string`.
* `interim_update` - (Optional) Type: `string`.
* `use_radius` - (Optional) Type: `bool`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_user_aaa.this 'home'
```
