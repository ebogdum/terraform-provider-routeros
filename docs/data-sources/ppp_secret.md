---
subcategory: "PPP"
page_title: "RouterOS: routeros_ppp_secret"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ppp_secret

Manages the RouterOS `/ppp/secret` menu.

## Example Usage

```terraform
data "routeros_ppp_secret" "secret_example" {
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
* `caller_id` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `ipv6` - (Optional) Type: `string`.
* `ipv6_routes` - (Optional) Type: `string`.
* `limit_bytes_in` - (Optional) Type: `string`.
* `limit_bytes_out` - (Optional) Type: `string`.
* `local_address` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_pppsec`.
* `password` - (Optional) Type: `string`. **Sensitive.**
* `profile` - (Optional) Type: `string`.
* `remote_address` - (Optional) Type: `string`.
* `remote_ipv6_prefix` - (Optional) Type: `string`.
* `routes` - (Optional) Type: `string`.
* `service` - (Optional) Type: `enum(any|async|pptp|pppoe|l2tp|ovpn, ...)`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`. **Marked sensitive**: this menu holds a secret, which RouterOS returns in the row like any other column, so an unprojected read puts it in your state file. Use `proplist` to name the columns you need.
* `last_caller_id` - Type: `string`.
* `last_disconnect_reason` - Type: `enum(|peer-request|hung-up|idle-timeout|session-timeout|reset, ...)`.
* `last_logged_out` - Type: `string`.

