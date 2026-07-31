---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_radius_incoming"
description: |-
  RouterOS resource.
---

# Resource: routeros_radius_incoming

Manages the RouterOS `/radius/incoming` menu.

## Example Usage

```terraform
resource "routeros_radius_incoming" "incoming_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # accept = false
  # port = "443"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `accept` - (Optional) Type: `bool`.
* `port` - (Optional) Type: `int`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_radius_incoming.this 'home'
```
