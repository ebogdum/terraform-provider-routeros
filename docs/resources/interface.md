---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface"
description: |-
  /interface is mostly read-only — interfaces are created by their specific subtypes (/interface/bridge, /interface/wireguard, etc.).
---

# Resource: routeros_interface

/interface is mostly read-only — interfaces are created by their specific subtypes (/interface/bridge, /interface/wireguard, etc.).

## Example Usage

```terraform
resource "routeros_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # fp_tx_rx_packet_rate = "replace-me"
  # fp_tx_rx_rate = "replace-me"
  # inactive = false
  # mtu = 1500
  # name = "tf-example"
  # nodefname = "replace-me"
  # notrunning = "replace-me"
  # passthrough = false
  # reset_traffic_counters = "replace-me"
  # slave = false
  # torch = "replace-me"
  # tx_rx_bytes = "replace-me"
  # tx_rx_drops = "replace-me"
  # tx_rx_errors = "replace-me"
  # tx_rx_packet_rate = "replace-me"
  # tx_rx_packets = "replace-me"
  # tx_rx_rate = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
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

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
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

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface.example '*3'

# Named router
terraform import routeros_interface.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface.example 'home/my-resource-name'
```
