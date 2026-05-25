---
subcategory: "Interfaces"
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

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `mac_address` - Type: `string`.

