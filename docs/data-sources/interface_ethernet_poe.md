---
page_title: "RouterOS: routeros_interface_ethernet_poe"
description: |-
  Requires PoE-capable ethernet port
---

# Data Source: routeros_interface_ethernet_poe

Requires PoE-capable ethernet port

## Example Usage

```terraform
data "routeros_interface_ethernet_poe" "poe_example" {
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
* `export` - (Optional) Type: `string`. export is displayed under   /interface ethernet   menu.
* `monitor` - (Optional) Type: `string`. Shows poe-out-status of a specified port, or all ports with   /interface ethernet poe monitor [find]   command.
* `power_cycle` - (Optional) Type: `string`. Disables PoE-Out power for a specified period of time.
* `print` - (Optional) Type: `string`. Prints PoE-Out related settings.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

