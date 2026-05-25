---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_user_manager_profile_limitation"
description: |-
  RouterOS resource.
---

# Resource: routeros_user_manager_profile_limitation

Manages the RouterOS `/user-manager/profile-limitation` menu.

## Example Usage

```terraform
resource "routeros_user_manager_profile_limitation" "profile_limitation_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_user_manager_profile_limitation.example '*3'

# Named router
terraform import routeros_user_manager_profile_limitation.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_user_manager_profile_limitation.example 'home/my-resource-name'
```
