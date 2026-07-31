---
subcategory: "IPv6"
page_title: "RouterOS: routeros_ipv6_neighbor"
description: |-
  IPv6 neighbor table — read-only on most devices.
---

# Resource: routeros_ipv6_neighbor

IPv6 neighbor table — read-only on most devices.

## Example Usage

```terraform
resource "routeros_ipv6_neighbor" "neighbor_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "10.99.0.1"
  # interface = "ether1"
  # mac_address = "10.99.0.0/24"
  # mac_ping = "replace-me"
  # mac_telnet = "replace-me"
  # make_static = "replace-me"
  # ping = "replace-me"
  # router = false
  # telnet = "replace-me"
  # torch = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Read-only) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address` - (Optional) Type: `string`.
* `bridge_port` - (Read-only) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dynamic` - (Read-only) Type: `bool`.
* `host_name` - (Read-only) Type: `string`.
* `interface` - (Optional) Type: `string`.
* `mac_address` - (Optional) Type: `string`.
* `mac_ping` - (Read-only) Type: `string`.
* `mac_telnet` - (Read-only) Type: `string`.
* `make_static` - (Read-only) Type: `string`.
* `ping` - (Read-only) Type: `string`.
* `router_ros` - (Optional) Type: `bool`.
* `status` - (Read-only) Type: `string`.
* `telnet` - (Read-only) Type: `string`.
* `torch` - (Read-only) Type: `string`.
* `vrf` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ipv6_neighbor.example '*3'

# Named router
terraform import routeros_ipv6_neighbor.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ipv6_neighbor.example 'home/my-resource-name'
```
