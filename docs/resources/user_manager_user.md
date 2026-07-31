---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_user_manager_user"
description: |-
  RouterOS resource.
---

# Resource: routeros_user_manager_user

Manages the RouterOS `/user-manager/user` menu.

## Example Usage

```terraform
resource "routeros_user_manager_user" "user_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `attributes` - (Optional) Type: `string`. RADIUS attributes returned on authentication, e.g. `Framed-IP-Address:10.0.0.5`.
* `caller_id` - (Optional) Type: `string`. RouterOS `caller-id`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `group` - (Optional) Type: `string`. User Manager group, e.g. `default`.
* `name` - (Optional) Type: `string`. Username.
* `otp_secret` - (Optional) Type: `string`. Base32 TOTP secret used for one-time-password authentication. **Sensitive.**
* `password` - (Optional) Type: `string`. User password. **Sensitive.**
* `shared_users` - (Optional) Type: `string`. Number of simultaneous sessions permitted for this user.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_user_manager_user.example '*3'

# Named router
terraform import routeros_user_manager_user.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_user_manager_user.example 'home/my-resource-name'
```
