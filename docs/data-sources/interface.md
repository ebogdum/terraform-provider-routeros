---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface"
description: |-
  /interface is mostly read-only — interfaces are created by their specific subtypes (/interface/bridge, /interface/wireguard, etc.).
---

# Data Source: routeros_interface

/interface is mostly read-only — interfaces are created by their specific subtypes (/interface/bridge, /interface/wireguard, etc.).

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
* `fp_tx_rx_packet_rate` - (Optional) Type: `string`.
* `fp_tx_rx_rate` - (Optional) Type: `string`.
* `inactive` - (Optional) Type: `bool`.
* `mtu` - (Optional) Type: `int`. Default: `1500`.
* `name` - (Optional) Type: `string`.
* `nodefname` - (Optional) Type: `string`.
* `notrunning` - (Optional) Type: `string`.
* `passthrough` - (Optional) Type: `bool`.
* `reset_traffic_counters` - (Optional) Type: `string`.
* `slave` - (Optional) Type: `bool`.
* `torch` - (Optional) Type: `string`.
* `tx_rx_bytes` - (Optional) Type: `string`.
* `tx_rx_drops` - (Optional) Type: `string`.
* `tx_rx_errors` - (Optional) Type: `string`.
* `tx_rx_packet_rate` - (Optional) Type: `string`.
* `tx_rx_packets` - (Optional) Type: `string`.
* `tx_rx_rate` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `actual_mtu` - Type: `int`.
* `answer_time` - Type: `string`.
* `caps` - Type: `int`.
* `default_name` - Type: `string`.
* `dynamic` - Type: `bool`.
* `fp_rps_drop` - Type: `int`.
* `fp_rx_byte` - Type: `int`.
* `fp_rx_packet` - Type: `int`.
* `fp_tx_byte` - Type: `int`.
* `fp_tx_packet` - Type: `int`.
* `l2_mtu` - Type: `int`.
* `last_link_down_time` - Type: `string`.
* `last_link_up_time` - Type: `string`.
* `link` - Type: `int`.
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
* `tx_queue_drops` - Type: `string`.
* `type` - Type: `int`.
* `vrf` - Type: `string`.

