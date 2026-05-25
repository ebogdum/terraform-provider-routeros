---
page_title: "RouterOS: routeros_interface_ovpn_client"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_ovpn_client

Manages the RouterOS `/interface/ovpn-client` menu.

## Example Usage

```terraform
data "routeros_interface_ovpn_client" "ovpn_client_example" {
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
* `certificate` - (Optional) Type: `string`.
* `cipher` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connect_to` - (Required) Type: `string`. Default: `127.0.0.1`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `mac_address` - (Optional) Type: `string`.
* `max_mtu` - (Optional) Type: `string`.
* `mode` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_ovpn_c`.
* `password` - (Optional) Type: `string`.
* `profile` - (Optional) Type: `string`.
* `user` - (Required) Type: `string`. Default: `tf_acc_user`.
* `verify_server_certificate` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

