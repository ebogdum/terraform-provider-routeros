---
page_title: "RouterOS: routeros_interface_vxlan"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_vxlan

Manages the RouterOS `/interface/vxlan` menu.

## Example Usage

```terraform
data "routeros_interface_vxlan" "vxlan_example" {
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
* `arp` - (Optional) Type: `string`.
* `arp_timeout` - (Optional) Type: `string`.
* `bridge` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Optional) Type: `string`.
* `local_address` - (Optional) Type: `string`.
* `mac_address` - (Optional) Type: `string`.
* `mtu` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_vxlan`.
* `port` - (Optional) Type: `string`.
* `ttl` - (Optional) Type: `string`.
* `vni` - (Required) Type: `string`. Default: `100`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

