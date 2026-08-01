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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `actual_mtu` - (Read-only) Type: `int`.
* `answer_time` - (Read-only) Type: `string`.
* `caps` - (Read-only) Type: `int`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default_name` - (Read-only) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `dynamic` - (Read-only) Type: `bool`.
* `fp_rps_drop` - (Read-only) Type: `int`.
* `fp_rx_byte` - (Read-only) Type: `int`.
* `fp_rx_packet` - (Read-only) Type: `int`.
* `fp_tx_byte` - (Read-only) Type: `int`.
* `fp_tx_packet` - (Read-only) Type: `int`.
* `fp_tx_rx_packet_rate` - (Read-only) Type: `string`.
* `fp_tx_rx_rate` - (Read-only) Type: `string`.
* `inactive` - (Read-only) Type: `bool`.
* `last_link_down_time` - (Read-only) Type: `string`.
* `last_link_up_time` - (Read-only) Type: `string`.
* `link` - (Read-only) Type: `int`.
* `link_downs` - (Read-only) Type: `int`.
* `mac_address` - (Read-only) Type: `string`.
* `mtu` - (Optional) Type: `string`. A number, or `auto`.
* `name` - (Optional) Type: `string`.
* `nodefname` - (Read-only) Type: `string`.
* `notrunning` - (Read-only) Type: `string`.
* `passthrough` - (Read-only) Type: `bool`.
* `reset_traffic_counters` - (Read-only) Type: `string`.
* `running` - (Read-only) Type: `bool`.
* `rx_byte` - (Read-only) Type: `int`.
* `rx_drop` - (Read-only) Type: `int`.
* `rx_error` - (Read-only) Type: `int`.
* `rx_packet` - (Read-only) Type: `int`.
* `slave` - (Read-only) Type: `bool`.
* `torch` - (Read-only) Type: `string`.
* `tx_byte` - (Read-only) Type: `int`.
* `tx_drop` - (Read-only) Type: `int`.
* `tx_error` - (Read-only) Type: `int`.
* `tx_packet` - (Read-only) Type: `int`.
* `tx_queue_drop` - (Read-only) Type: `int`.
* `tx_queue_drops` - (Read-only) Type: `string`.
* `tx_rx_bytes` - (Read-only) Type: `string`.
* `tx_rx_drops` - (Read-only) Type: `string`.
* `tx_rx_errors` - (Read-only) Type: `string`.
* `tx_rx_packet_rate` - (Read-only) Type: `string`.
* `tx_rx_packets` - (Read-only) Type: `string`.
* `tx_rx_rate` - (Read-only) Type: `string`.
* `type` - (Read-only) Type: `int`.
* `vrf` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
