---
page_title: "RouterOS: routeros_ipv6_neighbor"
description: |-
  IPv6 neighbor table -- read-only on most devices.
---

# Resource: routeros_ipv6_neighbor

IPv6 neighbor table -- read-only on most devices.

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
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Optional) Type: `ipv6`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Optional) Type: `string`.
* `mac_address` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `dynamic` - Type: `bool`.
* `status` - Type: `string`.
* `vrf` - Type: `string`.

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
