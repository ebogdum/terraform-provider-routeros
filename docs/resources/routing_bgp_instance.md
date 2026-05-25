---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_bgp_instance"
description: |-
  RouterOS resource.
---

# Resource: routeros_routing_bgp_instance

Manages the RouterOS `/routing/bgp/instance` menu.

## Example Usage

```terraform
resource "routeros_routing_bgp_instance" "instance_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # as = "replace-me"
  # cluster_id = "replace-me"
  # name = "tf-example"
  # router_id = "replace-me"
  # routing_table = "main"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `as` - (Optional) Type: `string`.
* `cluster_id` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `name` - (Optional) Type: `string`.
* `router_id` - (Optional) Type: `string`.
* `routing_table` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_bgp_instance.example '*3'

# Named router
terraform import routeros_routing_bgp_instance.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_bgp_instance.example 'home/my-resource-name'
```
