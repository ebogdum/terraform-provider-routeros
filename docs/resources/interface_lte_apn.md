---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_lte_apn"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_lte_apn

Manages the RouterOS `/interface/lte/apn` menu.

## Example Usage

```terraform
resource "routeros_interface_lte_apn" "apn_example" {
  # router = "my-router"  # which router to target; omit for the default
  apn = "internet"
  name = "tf-example"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # add_default_route = false
  # authentication = "replace-me"
  # default_route_distance = 0
  # ip_type = "replace-me"
  # use_network_apn = false
  # use_peer_dns = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `add_default_route` - (Optional) Type: `bool`.
* `apn` - (Required) Type: `string`.
* `authentication` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default` - (Read-only) Type: `bool`.
* `default_route_distance` - (Optional) Type: `int`.
* `ip_type` - (Optional) Type: `string`.
* `ipv6_interface` - (Optional) Type: `string`. RouterOS `ipv6-interface`.
* `name` - (Required) Type: `string`.
* `passthrough_interface` - (Optional) Type: `string`. RouterOS `passthrough-interface`.
* `passthrough_mac` - (Optional) Type: `string`. RouterOS `passthrough-mac`.
* `passthrough_subnet_size` - (Optional) Type: `string`. RouterOS `passthrough-subnet-size`.
* `password` - (Optional) Type: `string`. RouterOS `password`. **Sensitive.**
* `use_network_apn` - (Optional) Type: `bool`.
* `use_peer_dns` - (Optional) Type: `bool`.
* `user` - (Optional) Type: `string`. RouterOS `user`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_lte_apn.example '*3'

# Named router
terraform import routeros_interface_lte_apn.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_lte_apn.example 'home/my-resource-name'
```
