---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_mesh"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_mesh

Manages the RouterOS `/interface/mesh` menu.

## Example Usage

```terraform
resource "routeros_interface_mesh" "mesh_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # admin_mac_address = "10.99.0.0/24"
  # arp = "enabled"
  # arp_timeout = "1h"
  # default_hoplimit = 32
  # mesh_portal = false
  # mesh_traceroute = "replace-me"
  # mtu = 1500
  # prep_lifetime = "300"
  # preq_destination_only = true
  # preq_reply_and_forward = true
  # preq_retries = 2
  # preq_waiting_time = 4
  # rann_interval = "10"
  # rann_lifetime = "22"
  # rann_propagation_delay = 500
  # reoptimize_paths = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `admin_mac_address` - (Optional) Type: `string`.
* `arp` - (Optional) Type: `enum(disabled|enabled|proxy-arp|reply-only|local-proxy-arp)`. Default: `1`.
* `arp_timeout` - (Optional) Type: `duration`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default_hoplimit` - (Optional) Type: `int`. Default: `32`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `mesh_portal` - (Optional) Type: `bool`.
* `mesh_traceroute` - (Optional) Type: `string`.
* `mtu` - (Optional) Type: `int`. Default: `1500`.
* `prep_lifetime` - (Optional) Type: `duration`. Default: `300`.
* `preq_destination_only` - (Optional) Type: `bool`. Default: `1`.
* `preq_reply_and_forward` - (Optional) Type: `bool`. Default: `1`.
* `preq_retries` - (Optional) Type: `int`. Default: `2`.
* `preq_waiting_time` - (Optional) Type: `int`. Default: `4`.
* `rann_interval` - (Optional) Type: `duration`. Default: `10`.
* `rann_lifetime` - (Optional) Type: `duration`. Default: `22`.
* `rann_propagation_delay` - (Optional) Type: `int`. Default: `500`.
* `reoptimize_paths` - (Optional) Type: `bool`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `mac_address` - Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_mesh.example '*3'

# Named router
terraform import routeros_interface_mesh.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_mesh.example 'home/my-resource-name'
```
