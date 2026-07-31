---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_ovpn_client"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_ovpn_client

Manages the RouterOS `/interface/ovpn-client` menu.

## Example Usage

```terraform
resource "routeros_interface_ovpn_client" "ovpn_client_example" {
  # router = "my-router"  # which router to target; omit for the default
  connect_to = "127.0.0.1"
  name = "tf-example"
  user = "myuser"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # certificate = "replace-me"
  # cipher = "replace-me"
  # mac_address = "10.99.0.0/24"
  # max_mtu = "replace-me"
  # mode = "replace-me"
  # password = "REDACTED"
  # profile = "replace-me"
  # verify_server_certificate = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `add_default_route` - (Optional) Type: `string`. RouterOS `add-default-route`.
* `auth` - (Optional) Type: `string`. RouterOS `auth`.
* `certificate` - (Optional) Type: `string`.
* `cipher` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connect_to` - (Required) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `disconnect_notify` - (Optional) Type: `string`. RouterOS `disconnect-notify`.
* `mac_address` - (Optional) Type: `string`.
* `max_mtu` - (Optional) Type: `string`.
* `mode` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `password` - (Optional) Type: `string`. **Sensitive.**
* `port` - (Optional) Type: `string`. RouterOS `port`.
* `profile` - (Optional) Type: `string`.
* `protocol` - (Optional) Type: `string`. RouterOS `protocol`.
* `route_nopull` - (Optional) Type: `string`. RouterOS `route-nopull`.
* `tls_version` - (Optional) Type: `string`. RouterOS `tls-version`.
* `use_peer_dns` - (Optional) Type: `string`. RouterOS `use-peer-dns`.
* `user` - (Required) Type: `string`.
* `verify_server_certificate` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_ovpn_client.example '*3'

# Named router
terraform import routeros_interface_ovpn_client.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_ovpn_client.example 'home/my-resource-name'
```
