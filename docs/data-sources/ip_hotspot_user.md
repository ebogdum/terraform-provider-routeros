---
subcategory: "Hotspot"
page_title: "RouterOS: routeros_ip_hotspot_user"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_hotspot_user

Manages the RouterOS `/ip/hotspot/user` menu.

## Example Usage

```terraform
data "routeros_ip_hotspot_user" "user_example" {
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
* `address` - (Optional) Type: `ip`.
* `comment` - (Optional) Type: `string`.
* `def` - (Optional) Type: `bool`.
* `disabled` - (Optional) Type: `bool`.
* `email` - (Optional) Type: `string`.
* `limit_bytes_in` - (Optional) Type: `string`.
* `limit_bytes_out` - (Optional) Type: `string`.
* `limit_bytes_total` - (Optional) Type: `string`.
* `limit_uptime` - (Optional) Type: `duration`.
* `mac_address` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_user`.
* `nondef` - (Optional) Type: `string`.
* `nondefro` - (Optional) Type: `string`.
* `otp_secret` - (Optional) Type: `string`. **Sensitive.**
* `password` - (Optional) Type: `string`. **Sensitive.**
* `profile` - (Optional) Type: `string`.
* `reset_all_counters` - (Optional) Type: `string`.
* `reset_counters` - (Optional) Type: `string`.
* `routes` - (Optional) Type: `string`.
* `server` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`. **Marked sensitive**: this menu holds a secret, which RouterOS returns in the row like any other column, so an unprojected read puts it in your state file. Use `proplist` to name the columns you need.
* `bytes_in` - Type: `int`.
* `bytes_out` - Type: `int`.
* `default` - Type: `bool`.
* `dynamic` - Type: `bool`.
* `packets_in` - Type: `int`.
* `packets_out` - Type: `int`.
* `uptime` - Type: `duration`.

