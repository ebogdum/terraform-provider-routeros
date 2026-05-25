---
page_title: "RouterOS: routeros_interface_l2tp_client"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_l2tp_client

Manages the RouterOS `/interface/l2tp-client` menu.

## Example Usage

```terraform
data "routeros_interface_l2tp_client" "l2tp_client_example" {
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
* `allow` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connect_to` - (Required) Type: `string`. Default: `127.0.0.1`.
* `default_route_distance` - (Optional) Type: `string`.
* `dial_on_demand` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `ipsec_secret` - (Optional) Type: `string`.
* `keepalive_timeout` - (Optional) Type: `string`.
* `max_mru` - (Optional) Type: `string`.
* `max_mtu` - (Optional) Type: `string`.
* `mrru` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_l2tpc`.
* `password` - (Optional) Type: `string`.
* `profile` - (Optional) Type: `string`.
* `user` - (Required) Type: `string`. Default: `tf_acc_user`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

