---
subcategory: "WireGuard"
page_title: "RouterOS: routeros_interface_wireguard_peers"
description: |-
  Peer attached to a /interface/wireguard interface. Set the `interface`
---

# Data Source: routeros_interface_wireguard_peers

Peer attached to a /interface/wireguard interface. Set the `interface`
attribute to an existing WireGuard interface name.


## Example Usage

```terraform
data "routeros_interface_wireguard_peers" "peers_example" {
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
* `allowed_address` - (Optional) Type: `string`.
* `client_address` - (Optional) Type: `string`.
* `client_allowed_address` - (Optional) Type: `string`.
* `client_dns` - (Optional) Type: `string`.
* `client_endpoint` - (Optional) Type: `string`.
* `client_keepalive` - (Optional) Type: `duration`.
* `client_listen_port` - (Optional) Type: `int`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `endpoint` - (Optional) Type: `string`.
* `endpoint_address` - (Optional) Type: `string`.
* `endpoint_port` - (Optional) Type: `int`. Default: `0`.
* `interface` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `persistent_keepalive` - (Optional) Type: `duration`.
* `preshared_key` - (Optional) Type: `string`. **Sensitive.**
* `private_key` - (Optional) Type: `string`. **Sensitive.**
* `public_key` - (Optional) Type: `string`.
* `responder` - (Optional) Type: `bool`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`. **Marked sensitive**: this menu holds a secret, which RouterOS returns in the row like any other column, so an unprojected read puts it in your state file. Use `proplist` to name the columns you need.
* `client_config` - Type: `string`.
* `client_qr` - Type: `string`.
* `current_endpoint_address` - Type: `string`.
* `current_endpoint_port` - Type: `int`.
* `dynamic` - Type: `bool`.
* `last_handshake` - Type: `duration`.
* `rx` - Type: `string`.
* `tx` - Type: `string`.

