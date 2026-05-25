---
subcategory: "Special-login"
page_title: "RouterOS: routeros_special_login"
description: |-
  RouterOS resource.
---

# Resource: routeros_special_login

Manages the RouterOS `/special-login` menu.

## Example Usage

```terraform
resource "routeros_special_login" "special_login_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # channel = "replace-me"
  # port = "443"
  # user = "myuser"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `channel` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `port` - (Optional) Type: `string`.
* `user` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_special_login.example '*3'

# Named router
terraform import routeros_special_login.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_special_login.example 'home/my-resource-name'
```
