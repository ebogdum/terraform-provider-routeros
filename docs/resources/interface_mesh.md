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
  # arp = "enabled"
  # arp_timeout = "1h"
  # mesh_portal = false
  # mtu = 1500
  # reoptimize_paths = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `arp` - (Optional) Type: `enum(disabled|enabled|proxy-arp|reply-only|local-proxy-arp)`. Default: `1`.
* `arp_timeout` - (Optional) Type: `duration`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `mesh_portal` - (Optional) Type: `bool`.
* `mtu` - (Optional) Type: `int`. Default: `1500`.
* `reoptimize_paths` - (Optional) Type: `bool`.

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
