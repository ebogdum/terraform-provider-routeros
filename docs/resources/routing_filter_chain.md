---
page_title: "RouterOS: routeros_routing_filter_chain"
description: |-
  Discovered; required chain name must be unique
---

# Resource: routeros_routing_filter_chain

Discovered; required chain name must be unique

## Example Usage

```terraform
resource "routeros_routing_filter_chain" "chain_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # name = "tf-example"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `name` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_filter_chain.example '*3'

# Named router
terraform import routeros_routing_filter_chain.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_filter_chain.example 'home/my-resource-name'
```
