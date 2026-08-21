---
subcategory: "RIP"
page_title: "RouterOS: routeros_routing_rip_interface"
description: |-
  Discovered; needs rip instance
---

# Data Source: routeros_routing_rip_interface

Discovered; needs rip instance

## Example Usage

```terraform
data "routeros_routing_rip_interface" "interface_example" {
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
* `cost` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `instance` - (Optional) Type: `string`.
* `interfaces` - (Optional) Type: `string`.
* `key_chain` - (Optional) Type: `string`.
* `mode` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `password` - (Optional) Type: `string`. **Sensitive.**
* `poison_reverse` - (Optional) Type: `string`.
* `source_addresses` - (Optional) Type: `string`.
* `split_horizon` - (Optional) Type: `string`.
* `use_bfd` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`. **Marked sensitive**: this menu holds a secret, which RouterOS returns in the row like any other column, so an unprojected read puts it in your state file. Use `proplist` to name the columns you need.

