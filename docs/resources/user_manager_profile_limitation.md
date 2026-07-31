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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `from_time` - (Optional) Type: `string`. RouterOS `from-time`.
* `limitation` - (Optional) Type: `string`. RouterOS `limitation`.
* `profile` - (Optional) Type: `string`. RouterOS `profile`.
* `till_time` - (Optional) Type: `string`. RouterOS `till-time`.
* `weekdays` - (Optional) Type: `string`. RouterOS `weekdays`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
