---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_user"
description: |-
  User accounts. password is sensitive and not round-trippable (RouterOS scrubs it on read).
---

# Data Source: routeros_user

User accounts. password is sensitive and not round-trippable (RouterOS scrubs it on read).

## Example Usage

```terraform
data "routeros_user" "user_example" {
  # router   = "my-router"  # omit for the default router
  # filter   = { name = "some-name" }
  # proplist = ["name", "address"]
}
```

## Argument Reference

This data source supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to query.
* `filter` - (Optional) Map of field=value pairs to narrow the result set.
* `proplist` - (Optional) List of property names to project; smaller payload.
* `address` - (Optional) Type: `string`.
* `alias` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `group` - (Required) Type: `string`. Default: `read`.
* `inactivity_policy` - (Optional) Type: `string`.
* `inactivity_timeout` - (Optional) Type: `duration`.
* `name` - (Required) Type: `string`. Default: `tf_user`.
* `password` - (Required) Type: `string`. Default: `tf_pw`. **Sensitive.**
* `type` - (Optional) Type: `int`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `expired` - Type: `bool`.
* `last_logged_in` - Type: `string`.

