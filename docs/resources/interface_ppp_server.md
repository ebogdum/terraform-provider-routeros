---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_ppp_server"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_ppp_server

Manages the RouterOS `/interface/ppp-server` menu.

## Example Usage

```terraform
resource "routeros_interface_ppp_server" "ppp_server_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # authentication = "replace-me"
  # max_mru = "replace-me"
  # max_mtu = "replace-me"
  # mrru = "replace-me"
  # name = "tf-example"
  # profile = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `authentication` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `data_channel` - (Optional) Type: `string`. RouterOS `data-channel`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `max_mru` - (Optional) Type: `string`.
* `max_mtu` - (Optional) Type: `string`.
* `modem_init` - (Optional) Type: `string`. RouterOS `modem-init`.
* `mrru` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `null_modem` - (Optional) Type: `string`. RouterOS `null-modem`.
* `port` - (Optional) Type: `string`. RouterOS `port`.
* `profile` - (Optional) Type: `string`.
* `ring_count` - (Optional) Type: `string`. RouterOS `ring-count`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_ppp_server.example '*3'

# Named router
terraform import routeros_interface_ppp_server.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_ppp_server.example 'home/my-resource-name'
```
