---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_vrrp"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_vrrp

Manages the RouterOS `/interface/vrrp` menu.

## Example Usage

```terraform
data "routeros_interface_vrrp" "vrrp_example" {
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
* `authentication` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Required) Type: `string`.
* `interval` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_vrrp`.
* `password` - (Optional) Type: `string`.
* `priority` - (Optional) Type: `string`.
* `remote_address` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`. **Marked sensitive**: this menu holds a secret, which RouterOS returns in the row like any other column, so an unprojected read puts it in your state file. Use `proplist` to name the columns you need.

