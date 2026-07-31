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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `allowed_address` - (Optional) Type: `string`.
* `client_address` - (Optional) Type: `string`.
* `client_allowed_address` - (Optional) Type: `string`.
* `client_config` - (Read-only) Type: `string`.
* `client_dns` - (Optional) Type: `string`.
* `client_endpoint` - (Optional) Type: `string`.
* `client_keepalive` - (Optional) Type: `string`.
* `client_listen_port` - (Optional) Type: `int`.
* `client_qr` - (Read-only) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `current_endpoint_address` - (Read-only) Type: `string`.
* `current_endpoint_port` - (Read-only) Type: `int`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dynamic` - (Read-only) Type: `bool`.
* `endpoint` - (Read-only) Type: `string`.
* `endpoint_address` - (Optional) Type: `string`.
* `endpoint_port` - (Optional) Type: `int`.
* `interface` - (Optional) Type: `string`.
* `last_handshake` - (Read-only) Type: `string`.
* `name` - (Optional) Type: `string`.
* `persistent_keepalive` - (Optional) Type: `string`.
* `preshared_key` - (Optional) Type: `string`. **Sensitive.**
* `private_key` - (Optional) Type: `string`. **Sensitive.**
* `public_key` - (Optional) Type: `string`.
* `responder` - (Optional) Type: `bool`.
* `rx` - (Read-only) Type: `string`.
* `tx` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
