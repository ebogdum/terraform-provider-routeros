---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_ppp_client"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_ppp_client

Manages the RouterOS `/interface/ppp-client` menu.

## Example Usage

```terraform
resource "routeros_interface_ppp_client" "ppp_client_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # allow = "replace-me"
  # default_route_distance = "replace-me"
  # dial_on_demand = "replace-me"
  # keepalive_timeout = "replace-me"
  # max_mru = "replace-me"
  # max_mtu = "replace-me"
  # mrru = "replace-me"
  # name = "tf-example"
  # password = "REDACTED"
  # profile = "replace-me"
  # remote_address = "10.99.0.1"
  # user = "myuser"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `add_default_route` - (Optional) Type: `string`. RouterOS `add-default-route`.
* `allow` - (Optional) Type: `string`.
* `apn` - (Optional) Type: `string`. RouterOS `apn`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `data_channel` - (Optional) Type: `string`. RouterOS `data-channel`.
* `default_route_distance` - (Optional) Type: `string`.
* `dial_command` - (Optional) Type: `string`. RouterOS `dial-command`.
* `dial_on_demand` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `info_channel` - (Optional) Type: `string`. RouterOS `info-channel`.
* `keepalive_timeout` - (Optional) Type: `string`.
* `max_mru` - (Optional) Type: `string`.
* `max_mtu` - (Optional) Type: `string`.
* `modem_init` - (Optional) Type: `string`. RouterOS `modem-init`.
* `mrru` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `network_mode` - (Optional) Type: `string`. RouterOS `network-mode`.
* `null_modem` - (Optional) Type: `string`. RouterOS `null-modem`.
* `password` - (Optional) Type: `string`. **Sensitive.**
* `phone` - (Optional) Type: `string`. RouterOS `phone`.
* `pin` - (Optional) Type: `string`. RouterOS `pin`.
* `port` - (Optional) Type: `string`. RouterOS `port`.
* `profile` - (Optional) Type: `string`.
* `remote_address` - (Optional) Type: `string`.
* `use_peer_dns` - (Optional) Type: `string`. RouterOS `use-peer-dns`.
* `user` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_ppp_client.example '*3'

# Named router
terraform import routeros_interface_ppp_client.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_ppp_client.example 'home/my-resource-name'
```
