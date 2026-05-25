---
page_title: "RouterOS: routeros_ip_socks_users"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_socks_users

Manages the RouterOS `/ip/socks/users` menu.

## Example Usage

```terraform
data "routeros_ip_socks_users" "users_example" {
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
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `name` - (Required) Type: `string`. Default: `tf_acc_socksu`.
* `password` - (Required) Type: `string`. Default: `tf_acc_pw`. **Sensitive.**

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

