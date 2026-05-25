---
page_title: "RouterOS: routeros_interface_mesh"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_mesh

Manages the RouterOS `/interface/mesh` menu.

## Example Usage

```terraform
data "routeros_interface_mesh" "mesh_example" {
  # router   = "my-router"  # omit for the default router
  # filter   = { name = "some-name" }
  # proplist = ["name", "address"]
}
```

## Argument Reference

This data source supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to query.
* `filter` - (Optional) Map of field=value pairs to narrow the result set.
* `proplist` - (Optional) List of property names to project; smaller payload.
* `arp` - (Optional) Type: `enum(disabled|enabled|proxy-arp|reply-only|local-proxy-arp)`. Default: `1`.
* `arp_timeout` - (Optional) Type: `duration`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `mesh_portal` - (Optional) Type: `bool`.
* `mtu` - (Optional) Type: `int`. Default: `1500`.
* `reoptimize_paths` - (Optional) Type: `bool`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

