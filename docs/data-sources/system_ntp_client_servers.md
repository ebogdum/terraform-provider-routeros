---
subcategory: "NTP"
page_title: "RouterOS: routeros_system_ntp_client_servers"
description: |-
  NTP server list — accepts add but validator differs per ROS. Skipped from acc tests.
---

# Data Source: routeros_system_ntp_client_servers

NTP server list — accepts add but validator differs per ROS. Skipped from acc tests.

## Example Usage

```terraform
data "routeros_system_ntp_client_servers" "servers_example" {
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
* `auth_key` - (Optional) Type: `string`. **Sensitive.**
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `iburst` - (Optional) Type: `bool`. Default: `1`.
* `keys` - (Optional) Type: `string`.
* `max_poll` - (Optional) Type: `int`. Default: `10`.
* `min_poll` - (Optional) Type: `int`. Default: `6`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`. **Marked sensitive**: this menu holds a secret, which RouterOS returns in the row like any other column, so an unprojected read puts it in your state file. Use `proplist` to name the columns you need.
* `dynamic` - Type: `bool`.
* `resolved_address` - Type: `string`.

