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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `accounting` - (Optional) Type: `bool`.
* `default_group` - (Optional) Type: `string`.
* `exclude_groups` - (Optional) Type: `string`.
* `interim_update` - (Optional) Type: `duration`.
* `use_radius` - (Optional) Type: `bool`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_user_aaa.this 'home'
```
