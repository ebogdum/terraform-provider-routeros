---
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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
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
