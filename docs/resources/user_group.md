---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_user_group"
description: |-
  RouterOS resource.
---

# Resource: routeros_user_group

Manages the RouterOS `/user/group` menu.

## Example Usage

```terraform
resource "routeros_user_group" "group_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # policies = "replace-me"
  # policy = ["read"]
  # skin = "replace-me"
  # system = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `name` - (Required) Type: `string`. Default: `tf-acc-group`.
* `policies` - (Optional) Type: `string`.
* `policy` - (Optional) Type: `list`. Default: `read`.
* `skin` - (Optional) Type: `string`.
* `system` - (Optional) Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_user_group.example '*3'

# Named router
terraform import routeros_user_group.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_user_group.example 'home/my-resource-name'
```
