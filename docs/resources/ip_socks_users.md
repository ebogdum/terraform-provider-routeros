---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_socks_users"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_socks_users

Manages the RouterOS `/ip/socks/users` menu.

## Example Usage

```terraform
resource "routeros_ip_socks_users" "users_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "example"
  password = "REDACTED"

  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `name` - (Required) Type: `string`. Default: `tf_acc_socksu`.
* `password` - (Required) Type: `string`. Default: `tf_acc_pw`. **Sensitive.**

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_socks_users.example '*3'

# Named router
terraform import routeros_ip_socks_users.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_socks_users.example 'home/my-resource-name'
```
