---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_igmp_proxy_interface"
description: |-
  RouterOS resource.
---

# Resource: routeros_routing_igmp_proxy_interface

Manages the RouterOS `/routing/igmp-proxy/interface` menu.

## Example Usage

```terraform
resource "routeros_routing_igmp_proxy_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # alternative_subnets = "replace-me"
  # inactive = false
  # interface = "ether1"
  # threshold = 1
  # upstream = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `alternative_subnets` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `dynamic` - (Read-only) Type: `bool`.
* `inactive` - (Read-only) Type: `bool`.
* `interface` - (Optional) Type: `string`.
* `querier` - (Read-only) Type: `bool`.
* `rx_bytes` - (Read-only) Type: `string`.
* `rx_packets` - (Read-only) Type: `string`.
* `source_ip_address` - (Read-only) Type: `string`.
* `threshold` - (Optional) Type: `int`.
* `tx_bytes` - (Read-only) Type: `string`.
* `tx_packets` - (Read-only) Type: `string`.
* `upstream` - (Optional) Type: `bool`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_igmp_proxy_interface.example '*3'

# Named router
terraform import routeros_routing_igmp_proxy_interface.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_igmp_proxy_interface.example 'home/my-resource-name'
```
