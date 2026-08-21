---
subcategory: "System & misc"
page_title: "RouterOS: routeros_caps_man_interface"
description: |-
  CAPsMAN virtual interfaces are typically created automatically; manual creation collides with the master.
---

# Data Source: routeros_caps_man_interface

CAPsMAN virtual interfaces are typically created automatically; manual creation collides with the master.

## Example Usage

```terraform
data "routeros_caps_man_interface" "interface_example" {
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
* `arp_timeout` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `mac_address` - (Optional) Type: `mac`.
* `master_interface` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `radio_mac` - (Optional) Type: `mac`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`. **Marked sensitive**: this menu holds a secret, which RouterOS returns in the row like any other column, so an unprojected read puts it in your state file. Use `proplist` to name the columns you need.

