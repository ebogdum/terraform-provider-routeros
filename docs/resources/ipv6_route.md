---
subcategory: "IPv6"
page_title: "RouterOS: routeros_ipv6_route"
description: |-
  RouterOS resource.
---

# Resource: routeros_ipv6_route

Manages the RouterOS `/ipv6/route` menu.

## Example Usage

```terraform
resource "routeros_ipv6_route" "route_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # distance = 0
  # dst_address = "10.99.0.0/24"
  # gateway = "replace-me"
  # routing_table = "main"
  # scope = 0
  # target_scope = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `active` - (Read-only) Type: `bool`.
* `blackhole` - (Optional) Type: `string`. RouterOS `blackhole`.
* `check_gateway` - (Optional) Type: `string`. RouterOS `check-gateway`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connect` - (Read-only) Type: `bool`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `distance` - (Optional) Type: `int`.
* `dst_address` - (Optional) Type: `string`.
* `dynamic` - (Read-only) Type: `bool`.
* `gateway` - (Optional) Type: `string`.
* `immediate_gw` - (Read-only) Type: `string`.
* `inactive` - (Read-only) Type: `bool`.
* `pref_src` - (Optional) Type: `string`. RouterOS `pref-src`.
* `routing_table` - (Optional) Type: `string`.
* `scope` - (Optional) Type: `int`.
* `target_scope` - (Optional) Type: `int`.
* `vrf_interface` - (Optional) Type: `string`. RouterOS `vrf-interface`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ipv6_route.example '*3'

# Named router
terraform import routeros_ipv6_route.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ipv6_route.example 'home/my-resource-name'
```
