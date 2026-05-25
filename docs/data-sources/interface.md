---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface"
description: |-
  /interface is mostly read-only -- interfaces are created by their specific subtypes (/interface/bridge, /interface/wireguard, etc.).
---

# Data Source: routeros_interface

/interface is mostly read-only -- interfaces are created by their specific subtypes (/interface/bridge, /interface/wireguard, etc.).

## Example Usage

```terraform
data "routeros_interface" "interface_example" {
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
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `mtu` - (Optional) Type: `int`. Default: `1500`.
* `name` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `actual_mtu` - Type: `int`.
* `default_name` - Type: `string`.
* `fp_rps_drop` - Type: `int`.
* `fp_rx_byte` - Type: `int`.
* `fp_rx_packet` - Type: `int`.
* `fp_tx_byte` - Type: `int`.
* `fp_tx_packet` - Type: `int`.
* `last_link_up_time` - Type: `string`.
* `link_downs` - Type: `int`.
* `mac_address` - Type: `mac`.
* `running` - Type: `bool`.
* `rx_byte` - Type: `int`.
* `rx_drop` - Type: `int`.
* `rx_error` - Type: `int`.
* `rx_packet` - Type: `int`.
* `tx_byte` - Type: `int`.
* `tx_drop` - Type: `int`.
* `tx_error` - Type: `int`.
* `tx_packet` - Type: `int`.
* `tx_queue_drop` - Type: `int`.
* `type` - Type: `int`.
* `vrf` - Type: `string`.

