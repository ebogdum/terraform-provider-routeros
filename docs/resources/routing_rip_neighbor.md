---
page_title: "RouterOS: routeros_routing_rip_neighbor"
description: |-
  Discovered; needs rip neighbor fixture
---

# Resource: routeros_routing_rip_neighbor

Discovered; needs rip neighbor fixture

## Example Usage

```terraform
resource "routeros_routing_rip_neighbor" "neighbor_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "replace-me"
  # instance = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `instance` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_rip_neighbor.example '*3'

# Named router
terraform import routeros_routing_rip_neighbor.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_rip_neighbor.example 'home/my-resource-name'
```
