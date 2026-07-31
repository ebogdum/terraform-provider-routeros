---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_pppoe_client"
description: |-
  PPPoE client needs at least one interface in 'interfaces'. Skipped.
---

# Resource: routeros_interface_pppoe_client

PPPoE client needs at least one interface in 'interfaces'. Skipped.

## Example Usage

```terraform
resource "routeros_interface_pppoe_client" "pppoe_client_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # ac_name = "replace-me"
  # add_default_route = "replace-me"
  # allow = "replace-me"
  # default_route_distance = "replace-me"
  # dial_on_demand = "replace-me"
  # interface = "ether1"
  # keepalive_timeout = "replace-me"
  # max_mru = "replace-me"
  # max_mtu = "replace-me"
  # mrru = "replace-me"
  # name = "tf-example"
  # password = "REDACTED"
  # profile = "replace-me"
  # service_name = "replace-me"
  # use_peer_dns = "replace-me"
  # user = "myuser"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `ac_name` - (Optional) Type: `string`.
* `add_default_route` - (Optional) Type: `string`.
* `allow` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default_route_distance` - (Optional) Type: `string`.
* `dial_on_demand` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `host_uniq` - (Optional) Type: `string`. RouterOS `host-uniq`.
* `interface` - (Optional) Type: `string`.
* `keepalive_timeout` - (Optional) Type: `string`.
* `max_mru` - (Optional) Type: `string`.
* `max_mtu` - (Optional) Type: `string`.
* `mrru` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `password` - (Optional) Type: `string`. **Sensitive.**
* `profile` - (Optional) Type: `string`.
* `service_name` - (Optional) Type: `string`.
* `use_peer_dns` - (Optional) Type: `string`.
* `user` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_pppoe_client.example '*3'

# Named router
terraform import routeros_interface_pppoe_client.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_pppoe_client.example 'home/my-resource-name'
```
