---
subcategory: "WireGuard"
page_title: "RouterOS: routeros_interface_wireguard_peers"
description: |-
  Peer attached to a /interface/wireguard interface. Set the `interface`
---

# Resource: routeros_interface_wireguard_peers

Peer attached to a /interface/wireguard interface. Set the `interface`
attribute to an existing WireGuard interface name.


## Example Usage

```terraform
resource "routeros_interface_wireguard_peers" "peers_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # allowed_address = "10.99.0.0/24"
  # client_address = "10.99.0.0/24"
  # client_allowed_address = "10.99.0.0/24"
  # client_dns = "replace-me"
  # client_endpoint = "replace-me"
  # client_keepalive = "1h"
  # client_listen_port = "443"
  # endpoint = "replace-me"
  # endpoint_address = "10.99.0.0/24"
  # endpoint_port = "443"
  # interface = "ether1"
  # name = "tf-example"
  # persistent_keepalive = "1h"
  # preshared_key = "REDACTED"
  # private_key = "REDACTED"
  # public_key = "REDACTED"
  # responder = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
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

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `client_config` - Type: `string`.
* `client_qr` - Type: `string`.
* `current_endpoint_address` - Type: `string`.
* `current_endpoint_port` - Type: `int`.
* `dynamic` - Type: `bool`.
* `last_handshake` - Type: `duration`.
* `rx` - Type: `string`.
* `tx` - Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_wireguard_peers.example '*3'

# Named router
terraform import routeros_interface_wireguard_peers.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_wireguard_peers.example 'home/my-resource-name'
```
