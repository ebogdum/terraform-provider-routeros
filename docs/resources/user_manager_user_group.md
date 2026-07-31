---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_user_manager_user_group"
description: |-
  RouterOS resource.
---

# Resource: routeros_user_manager_user_group

Manages the RouterOS `/user-manager/user/group` menu.

## Example Usage

```terraform
resource "routeros_user_manager_user_group" "group_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # attributes = "replace-me"
  # default = "replace-me"
  # default_name = "replace-me"
  # inner_auths = "replace-me"
  # name = "tf-example"
  # outer_auths = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `attributes` - (Optional) Type: `string`.
* `default` - (Optional) Type: `string`.
* `default_name` - (Optional) Type: `string`.
* `inner_auths` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `outer_auths` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_user_manager_user_group.example '*3'

# Named router
terraform import routeros_user_manager_user_group.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_user_manager_user_group.example 'home/my-resource-name'
```
