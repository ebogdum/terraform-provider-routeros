---
subcategory: "RIP"
page_title: "RouterOS: routeros_routing_rip_instance"
description: |-
  RouterOS resource.
---

# Resource: routeros_routing_rip_instance

Manages the RouterOS `/routing/rip/instance` menu.

## Example Usage

```terraform
resource "routeros_routing_rip_instance" "instance_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # afi = "replace-me"
  # input_filter = "replace-me"
  # name = "tf-example"
  # originate_default = "replace-me"
  # output_filter = "replace-me"
  # redistribute = "replace-me"
  # route_gc_timeout = "replace-me"
  # route_timeout = "replace-me"
  # routing_table = "main"
  # select_output_filter = "replace-me"
  # update_interval = "replace-me"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `afi` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `input_filter` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `originate_default` - (Optional) Type: `string`.
* `output_filter` - (Optional) Type: `string`.
* `redistribute` - (Optional) Type: `string`.
* `route_gc_timeout` - (Optional) Type: `string`.
* `route_timeout` - (Optional) Type: `string`.
* `routing_table` - (Optional) Type: `string`.
* `select_output_filter` - (Optional) Type: `string`.
* `update_interval` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_rip_instance.example '*3'

# Named router
terraform import routeros_routing_rip_instance.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_rip_instance.example 'home/my-resource-name'
```
