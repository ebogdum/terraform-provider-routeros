---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_ethernet"
description: |-
  /interface/ethernet entries are auto-created from the physical NICs; can't be added via TF.
---

# Resource: routeros_interface_ethernet

/interface/ethernet entries are auto-created from the physical NICs; can't be added via TF.

## Example Usage

```terraform
resource "routeros_interface_ethernet" "ethernet_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # advertise = []
  # arp = "disabled"
  # arp_timeout = "1h"
  # autoneg = false
  # blink = "replace-me"
  # cable_settings = "replace-me"
  # cable_test = "replace-me"
  # combo_mode = "auto"
  # disable_running_check = false
  # disable_time = "300"
  # extrastats = "replace-me"
  # fec_mode = "off"
  # flowcntrl = "replace-me"
  # ignore_rx_los = false
  # loop_protect = "default"
  # loop_protect_disable_time = "1h"
  # loop_protect_send_interval = "1h"
  # mac_address = "10.99.0.0/24"
  # mtu = 1500
  # name = "tf-example"
  # noautoneg = "replace-me"
  # non_mgmt = "replace-me"
  # orig_mac_address = "10.99.0.0/24"
  # passthrough_interface = "replace-me"
  # poe_out = "off"
  # poe_priority = 0
  # poe_voltage = "auto"
  # poe = "replace-me"
  # poeping = "replace-me"
  # power_cycle = "replace-me"
  # power_cycle_interval = "1h"
  # power_cycle_ping_address = "10.99.0.0/24"
  # power_cycle_ping_enabled = false
  # power_cycle_ping_timeout = "1h"
  # qstats = "replace-me"
  # rate_select = "low"
  # reset_counters = "replace-me"
  # reset_mac_address = "10.99.0.0/24"
  # rx_flow_control = "off"
  # send_interval = "5"
  # sfp = false
  # sfp_shutdown_temperature = 0
  # speed = "10m-baset-half"
  # tx_flow_control = "off"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `advertise` - (Optional) Type: `list`.
* `advertising` - (Read-only) Type: `string`.
* `arp` - (Optional) Type: `string`.
* `arp_timeout` - (Optional) Type: `string`.
* `auto_negotiation` - (Optional) Type: `bool`. Whether the port negotiates link speed and duplex with its peer. Disable it only when forcing `speed`/`full_duplex` on both ends.
* `autoneg` - (Optional) Type: `bool`.
* `bandwidth` - (Optional) Type: `string`. Rx/tx rate limit in `<rx>/<tx>` form, e.g. `unlimited/unlimited` (the default) or `100M/100M`.
* `blink` - (Optional) Type: `string`.
* `cable_assembly_link_length` - (Read-only) Type: `string`.
* `cable_settings` - (Optional) Type: `string`.
* `cable_test` - (Optional) Type: `string`.
* `cmis_module_state` - (Read-only) Type: `string`.
* `cmis_revision` - (Read-only) Type: `string`.
* `combo` - (Read-only) Type: `int`.
* `combo_mode` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connector_type` - (Read-only) Type: `string`.
* `copper_active_om4_link_length` - (Read-only) Type: `int`.
* `default_name` - (Read-only) Type: `string`.
* `disable_running_check` - (Optional) Type: `bool`.
* `disable_time` - (Read-only) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `encoding` - (Read-only) Type: `string`.
* `extrastats` - (Optional) Type: `string`.
* `fec` - (Read-only) Type: `int`.
* `fec_mode` - (Optional) Type: `string`.
* `flowcntrl` - (Optional) Type: `string`.
* `flowcontrol` - (Read-only) Type: `int`.
* `full_duplex` - (Read-only) Type: `string`.
* `hastxqueuestats` - (Read-only) Type: `bool`.
* `ignore_rx_los` - (Optional) Type: `bool`.
* `l2mtu` - (Optional) Type: `string`. RouterOS `l2mtu`.
* `link_partner_advertising` - (Read-only) Type: `string`.
* `loop_protect` - (Optional) Type: `string`.
* `loop_protect_disable_time` - (Optional) Type: `string`.
* `loop_protect_send_interval` - (Optional) Type: `string`.
* `loop_protect_status` - (Read-only) Type: `string`.
* `mac_address` - (Optional) Type: `string`. MAC address to be mapped to
* `manufacturing_date` - (Read-only) Type: `string`.
* `max_l2_mtu` - (Read-only) Type: `int`.
* `max_power` - (Read-only) Type: `string`.
* `mdix_enable` - (Optional) Type: `string`. RouterOS `mdix-enable`.
* `module_present` - (Read-only) Type: `bool`.
* `mtu` - (Optional) Type: `int`.
* `name` - (Optional) Type: `string`.
* `noautoneg` - (Optional) Type: `string`.
* `non_mgmt` - (Read-only) Type: `string`.
* `om1_link_length` - (Read-only) Type: `int`.
* `om2_link_length` - (Read-only) Type: `int`.
* `om3_link_length` - (Read-only) Type: `int`.
* `om4_link_length` - (Read-only) Type: `int`.
* `om5_link_length` - (Read-only) Type: `int`.
* `orig_mac_address` - (Optional) Type: `string`.
* `passthrough_interface` - (Read-only) Type: `string`.
* `pcie_passthrough` - (Read-only) Type: `int`.
* `poe` - (Optional) Type: `string`.
* `poe_out` - (Optional) Type: `string`.
* `poe_out_current` - (Optional) Type: `int`.
* `poe_out_power` - (Optional) Type: `string`.
* `poe_out_status` - (Optional) Type: `string`.
* `poe_out_voltage` - (Optional) Type: `string`.
* `poe_priority` - (Optional) Type: `int`.
* `poe_v` - (Read-only) Type: `bool`.
* `poe_voltage` - (Optional) Type: `string`.
* `poecurr` - (Read-only) Type: `int`.
* `poeping` - (Optional) Type: `string`.
* `poepower` - (Read-only) Type: `int`.
* `poevolt` - (Read-only) Type: `int`.
* `power_class` - (Read-only) Type: `int`.
* `power_cycle` - (Read-only) Type: `string`.
* `power_cycle_after` - (Read-only) Type: `string`.
* `power_cycle_host_alive` - (Read-only) Type: `bool`.
* `power_cycle_interval` - (Optional) Type: `string`.
* `power_cycle_ping_address` - (Read-only) Type: `string`.
* `power_cycle_ping_enabled` - (Optional) Type: `bool`.
* `power_cycle_ping_timeout` - (Read-only) Type: `string`.
* `qstats` - (Optional) Type: `string`.
* `rate` - (Read-only) Type: `string`.
* `rate_select` - (Optional) Type: `string`.
* `reset_counters` - (Read-only) Type: `string`.
* `reset_mac_address` - (Read-only) Type: `string`.
* `running` - (Read-only) Type: `bool`.
* `rx_align_error` - (Read-only) Type: `string`.
* `rx_broadcast` - (Read-only) Type: `int`.
* `rx_bytes` - (Read-only) Type: `int`.
* `rx_carrier_error` - (Read-only) Type: `string`.
* `rx_code_error` - (Read-only) Type: `string`.
* `rx_control` - (Read-only) Type: `string`.
* `rx_drop` - (Read-only) Type: `string`.
* `rx_error_events` - (Read-only) Type: `string`.
* `rx_fcs_error` - (Read-only) Type: `string`.
* `rx_flow_control` - (Optional) Type: `string`.
* `rx_fragment` - (Read-only) Type: `string`.
* `rx_jabber` - (Read-only) Type: `string`.
* `rx_length_error` - (Read-only) Type: `string`.
* `rx_loss` - (Read-only) Type: `bool`.
* `rx_multicast` - (Read-only) Type: `int`.
* `rx_overflow` - (Read-only) Type: `string`.
* `rx_packet` - (Read-only) Type: `int`.
* `rx_pause` - (Read-only) Type: `string`.
* `rx_power` - (Read-only) Type: `string`.
* `rx_too_long` - (Read-only) Type: `string`.
* `rx_too_short` - (Read-only) Type: `string`.
* `rx_unicast` - (Read-only) Type: `string`.
* `rx_unknown_op` - (Read-only) Type: `string`.
* `send_interval` - (Read-only) Type: `string`.
* `sfp` - (Optional) Type: `bool`.
* `sfp_ignore_rx_los` - (Optional) Type: `string`. RouterOS `sfp-ignore-rx-los`.
* `sfp_rate_select` - (Optional) Type: `string`. RouterOS `sfp-rate-select`.
* `sfp_shutdown_temperature` - (Optional) Type: `int`.
* `sfp_supported` - (Read-only) Type: `string`.
* `sfprate` - (Read-only) Type: `int`.
* `sfpshutdown` - (Read-only) Type: `bool`.
* `sm_link_length` - (Read-only) Type: `string`.
* `speed` - (Optional) Type: `string`.
* `status` - (Read-only) Type: `string`.
* `supply_voltage` - (Read-only) Type: `string`.
* `supported` - (Read-only) Type: `string`.
* `temperature` - (Read-only) Type: `string`.
* `tx_bias_current` - (Read-only) Type: `int`.
* `tx_broadcast` - (Read-only) Type: `int`.
* `tx_bytes` - (Read-only) Type: `int`.
* `tx_collision` - (Read-only) Type: `string`.
* `tx_control` - (Read-only) Type: `string`.
* `tx_deferred` - (Read-only) Type: `string`.
* `tx_drop` - (Read-only) Type: `string`.
* `tx_excessive_collision` - (Read-only) Type: `string`.
* `tx_excessive_deferred` - (Read-only) Type: `string`.
* `tx_fault` - (Read-only) Type: `bool`.
* `tx_fcs_error` - (Read-only) Type: `string`.
* `tx_flow_control` - (Optional) Type: `string`.
* `tx_fragment` - (Read-only) Type: `string`.
* `tx_jabber` - (Read-only) Type: `string`.
* `tx_late_collision` - (Read-only) Type: `string`.
* `tx_multicast` - (Read-only) Type: `int`.
* `tx_multiple_collision` - (Read-only) Type: `string`.
* `tx_packet` - (Read-only) Type: `int`.
* `tx_pause` - (Read-only) Type: `string`.
* `tx_pause_honorred` - (Read-only) Type: `string`.
* `tx_power` - (Read-only) Type: `string`.
* `tx_rx_1024_1518` - (Read-only) Type: `string`.
* `tx_rx_1024_max` - (Read-only) Type: `string`.
* `tx_rx_128_255` - (Read-only) Type: `string`.
* `tx_rx_1519_max` - (Read-only) Type: `string`.
* `tx_rx_256_511` - (Read-only) Type: `string`.
* `tx_rx_512_1023` - (Read-only) Type: `string`.
* `tx_rx_64` - (Read-only) Type: `string`.
* `tx_rx_65_127` - (Read-only) Type: `string`.
* `tx_rx_bytes` - (Read-only) Type: `string`.
* `tx_rx_packets` - (Read-only) Type: `string`.
* `tx_single_collision` - (Read-only) Type: `string`.
* `tx_too_short` - (Read-only) Type: `string`.
* `tx_total_collision` - (Read-only) Type: `string`.
* `tx_underrun` - (Read-only) Type: `string`.
* `tx_unicast` - (Read-only) Type: `string`.
* `vendor_name` - (Read-only) Type: `string`.
* `vendor_part_number` - (Read-only) Type: `string`.
* `vendor_revision` - (Read-only) Type: `string`.
* `vendor_serial` - (Read-only) Type: `string`.
* `wavelength` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_ethernet.example '*3'

# Named router
terraform import routeros_interface_ethernet.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_ethernet.example 'home/my-resource-name'
```
