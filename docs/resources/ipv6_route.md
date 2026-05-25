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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `distance` - (Optional) Type: `int`.
* `dst_address` - (Optional) Type: `cidr`.
* `gateway` - (Optional) Type: `string`.
* `routing_table` - (Optional) Type: `string`.
* `scope` - (Optional) Type: `int`.
* `target_scope` - (Optional) Type: `int`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `active` - Type: `bool`.
* `connect` - Type: `bool`.
* `dynamic` - Type: `bool`.
* `immediate_gw` - Type: `string`.
* `inactive` - Type: `bool`.

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
