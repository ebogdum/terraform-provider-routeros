---
subcategory: "Users & RADIUS"
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
  name = "tf-example"
  password = "REDACTED"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "replace-me"
  # alias = "replace-me"
  # inactivity_policy = "replace-me"
  # inactivity_timeout = "1h"
  # type = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address` - (Optional) Type: `string`.
* `alias` - (Read-only) Type: `string`.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `expired` - (Read-only) Type: `bool`.
* `group` - (Required) Type: `string`.
* `inactivity_policy` - (Optional) Type: `string`.
* `inactivity_timeout` - (Optional) Type: `string`.
* `last_logged_in` - (Read-only) Type: `string`.
* `lockout_ack` - (Optional) Type: `bool`. Acknowledge that this rule may sever management traffic (required for unconditional input/forward drop/reject/tarpit rules with no match).
* `name` - (Required) Type: `string`.
* `password` - (Required) Type: `string`. **Sensitive.**
* `type` - (Read-only) Type: `int`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
