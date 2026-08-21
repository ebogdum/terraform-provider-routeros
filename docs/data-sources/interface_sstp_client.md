---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_sstp_client"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_sstp_client

Manages the RouterOS `/interface/sstp-client` menu.

## Example Usage

```terraform
data "routeros_interface_sstp_client" "sstp_client_example" {
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
* `authentication` - (Optional) Type: `string`.
* `certificate` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connect_to` - (Required) Type: `string`. Default: `127.0.0.1`.
* `default_route_distance` - (Optional) Type: `string`.
* `dial_on_demand` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `keepalive_timeout` - (Optional) Type: `string`.
* `max_mru` - (Optional) Type: `string`.
* `max_mtu` - (Optional) Type: `string`.
* `mrru` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_sstpc`.
* `password` - (Optional) Type: `string`.
* `profile` - (Optional) Type: `string`.
* `user` - (Required) Type: `string`. Default: `tf_acc_user`.
* `verify_server_certificate` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`. **Marked sensitive**: this menu holds a secret, which RouterOS returns in the row like any other column, so an unprojected read puts it in your state file. Use `proplist` to name the columns you need.

