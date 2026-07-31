---
subcategory: "Routing"
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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `dynamic` - (Read-only) Type: `bool`.
* `invalid` - (Read-only) Type: `bool`.
* `name` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
