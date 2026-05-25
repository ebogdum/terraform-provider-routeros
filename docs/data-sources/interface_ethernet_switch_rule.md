---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_ethernet_switch_rule"
description: |-
  Requires switch chip presence
---

# Data Source: routeros_interface_ethernet_switch_rule

Requires switch chip presence

## Example Usage

```terraform
data "routeros_interface_ethernet_switch_rule" "rule_example" {
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
* `qos_hw_offloading` - (Optional) Type: `string`. Allows enabling QoS for the given switch chip (if the latter supports QoS). New generation devices force qos-hw-offloading=yes at all times.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

