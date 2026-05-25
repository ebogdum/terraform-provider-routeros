---
page_title: "RouterOS: routeros_caps_man_datapath"
description: |-
  RouterOS resource.
---

# Data Source: routeros_caps_man_datapath

Manages the RouterOS `/caps-man/datapath` menu.

## Example Usage

```terraform
data "routeros_caps_man_datapath" "datapath_example" {
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
* `bridge` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `l2mtu` - (Optional) Type: `string`.
* `mtu` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `ros_audit_20260523213235_5`.
* `vlan_id` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

