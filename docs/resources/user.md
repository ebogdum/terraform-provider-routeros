---
subcategory: "Users"
page_title: "RouterOS: routeros_user"
description: |-
  User accounts. password is sensitive and not round-trippable (RouterOS scrubs it on read).
---

# Resource: routeros_user

User accounts. password is sensitive and not round-trippable (RouterOS scrubs it on read).

## Example Usage

```terraform
resource "routeros_user" "user_example" {
  # router = "my-router"  # which router to target; omit for the default
  group = "read"
  name = "example"
  password = "REDACTED"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "10.99.0.1"
  # inactivity_policy = "replace-me"
  # inactivity_timeout = "1h"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `group` - (Required) Type: `string`. Default: `read`.
* `inactivity_policy` - (Optional) Type: `string`.
* `inactivity_timeout` - (Optional) Type: `duration`.
* `name` - (Required) Type: `string`. Default: `tf_user`.
* `password` - (Required) Type: `string`. Default: `tf_pw`. **Sensitive.**

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `expired` - Type: `bool`.
* `last_logged_in` - Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_user.example '*3'

# Named router
terraform import routeros_user.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_user.example 'home/my-resource-name'
```
