---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_user_manager_profile"
description: |-
  RouterOS resource.
---

# Resource: routeros_user_manager_profile

Manages the RouterOS `/user-manager/profile` menu.

## Example Usage

```terraform
resource "routeros_user_manager_profile" "profile_example" {
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
terraform import routeros_user_manager_profile.example '*3'

# Named router
terraform import routeros_user_manager_profile.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_user_manager_profile.example 'home/my-resource-name'
```
