---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_ovpn_server"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_ovpn_server

Manages the RouterOS `/interface/ovpn-server` menu.

## Example Usage

```terraform
resource "routeros_interface_ovpn_server" "ovpn_server_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"
  user = "myuser"

  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `name` - (Required) Type: `string`.
* `user` - (Required) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_ovpn_server.example '*3'

# Named router
terraform import routeros_interface_ovpn_server.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_ovpn_server.example 'home/my-resource-name'
```
