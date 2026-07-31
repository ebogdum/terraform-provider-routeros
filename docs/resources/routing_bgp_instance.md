---
subcategory: "BGP"
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
  # ignore_as_path = "replace-me"
  # name = "tf-example"
  # router_id = "replace-me"
  # routing_table = "main"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `as` - (Optional) Type: `string`.
* `cluster_id` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `ignore_as_path` - (Read-only) Type: `string`.
* `ignore_as_path_len` - (Optional) Type: `string`. RouterOS `ignore-as-path-len`.
* `invalid` - (Read-only) Type: `bool`.
* `multipath` - (Optional) Type: `string`. RouterOS `multipath`.
* `name` - (Optional) Type: `string`.
* `router_id` - (Optional) Type: `string`.
* `routing_table` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
