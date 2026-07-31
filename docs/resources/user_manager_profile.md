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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `name` - (Optional) Type: `string`. RouterOS `name`.
* `name_for_users` - (Optional) Type: `string`. RouterOS `name-for-users`.
* `override_shared_users` - (Optional) Type: `string`. RouterOS `override-shared-users`.
* `price` - (Optional) Type: `string`. RouterOS `price`.
* `starts_when` - (Optional) Type: `string`. RouterOS `starts-when`.
* `validity` - (Optional) Type: `string`. RouterOS `validity`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
