---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_ethernet"
description: |-
  /interface/ethernet entries are auto-created from the physical NICs; can't be added via TF.
---

# Data Source: routeros_interface_ethernet

/interface/ethernet entries are auto-created from the physical NICs; can't be added via TF.

## Example Usage

```terraform
data "routeros_interface_ethernet" "ethernet_example" {
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
* `advertise` - (Optional) Type: `list`.
* `arp` - (Optional) Type: `enum(disabled|enabled|proxy-arp|reply-only|local-proxy-arp)`.
* `arp_timeout` - (Optional) Type: `duration`.
* `autoneg` - (Optional) Type: `bool`.
* `blink` - (Optional) Type: `string`.
* `cable_settings` - (Optional) Type: `string`.
* `cable_test` - (Optional) Type: `string`.
* `combo_mode` - (Optional) Type: `enum(auto|copper|sfp)`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disable_running_check` - (Optional) Type: `bool`.
* `disable_time` - (Optional) Type: `duration`. Default: `300`.
* `disabled` - (Optional) Type: `bool`.
* `extrastats` - (Optional) Type: `string`.
* `fec_mode` - (Optional) Type: `enum(off|auto|fec74|fec91)`.
* `flowcntrl` - (Optional) Type: `string`.
* `ignore_rx_los` - (Optional) Type: `bool`.
* `l2_mtu` - (Optional) Type: `int`.
* `loop_protect` - (Optional) Type: `enum(default|off|on)`.
* `loop_protect_disable_time` - (Optional) Type: `duration`.
* `loop_protect_send_interval` - (Optional) Type: `duration`.
* `mac_address` - (Optional) Type: `mac`. MAC address to be mapped to.
* `mtu` - (Optional) Type: `int`. Default: `1500`.
* `name` - (Optional) Type: `string`.
* `noautoneg` - (Optional) Type: `string`.
* `non_mgmt` - (Optional) Type: `string`.
* `orig_mac_address` - (Optional) Type: `mac`.
* `passthrough_interface` - (Optional) Type: `string`.
* `po_e_out` - (Optional) Type: `enum(off|auto-on|forced-on)`.
* `po_e_priority` - (Optional) Type: `int`.
* `po_e_voltage` - (Optional) Type: `enum(auto|low|high)`.
* `poe` - (Optional) Type: `string`.
* `poeping` - (Optional) Type: `string`.
* `power_cycle` - (Optional) Type: `string`.
* `power_cycle_interval` - (Optional) Type: `duration`.
* `power_cycle_ping_address` - (Optional) Type: `string`.
* `power_cycle_ping_enabled` - (Optional) Type: `bool`.
* `power_cycle_ping_timeout` - (Optional) Type: `duration`.
* `qstats` - (Optional) Type: `string`.
* `rate_select` - (Optional) Type: `enum(low|high)`.
* `reset_counters` - (Optional) Type: `string`.
* `reset_mac_address` - (Optional) Type: `string`.
* `rx_flow_control` - (Optional) Type: `enum(off|on|auto)`.
* `send_interval` - (Optional) Type: `duration`. Default: `5`.
* `sfp` - (Optional) Type: `bool`.
* `sfp_shutdown_temperature` - (Optional) Type: `int`.
* `speed` - (Optional) Type: `enum(10m-baset-half|10m-baset-full|100m-baset-half|100m-baset-full|1g-baset-half|1g-baset-full, ...)`.
* `tx_flow_control` - (Optional) Type: `enum(off|on|auto)`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `advertising` - Type: `string`.
* `auto_negotiation` - Type: `enum(incomplete|done|no-negotiation|failed|restarted|disabled, ...)`.
* `cable_assembly_link_length` - Type: `string`.
* `cmis_module_state` - Type: `enum(low-power|power-up|ready|power-down|fault)`.
* `cmis_revision` - Type: `string`.
* `combo` - Type: `int`.
* `connector_type` - Type: `enum(unknown|sc|lc|optical-pigtail|multifiber-parallel-optic-1x12|copper-pigtail, ...)`.
* `copper_active_om4_link_length` - Type: `int`.
* `default_name` - Type: `string`.
* `encoding` - Type: `enum(unspecified|8b/10b|4b/5b|nrz|manchester|sonet, ...)`.
* `fec` - Type: `int`.
* `flowcontrol` - Type: `int`.
* `full_duplex` - Type: `string`.
* `hastxqueuestats` - Type: `bool`.
* `link_partner_advertising` - Type: `string`.
* `loop_protect_status` - Type: `string`.
* `manufacturing_date` - Type: `string`.
* `max_l2_mtu` - Type: `int`.
* `max_power` - Type: `string`.
* `module_present` - Type: `bool`.
* `om1_link_length` - Type: `int`.
* `om2_link_length` - Type: `int`.
* `om3_link_length` - Type: `int`.
* `om4_link_length` - Type: `int`.
* `om5_link_length` - Type: `int`.
* `pcie_passthrough` - Type: `int`.
* `po_e_out_current` - Type: `int`.
* `po_e_out_power` - Type: `string`.
* `po_e_out_status` - Type: `enum(|disabled|waiting-for-load|powered-on|overload|short-circuit, ...)`.
* `po_e_out_voltage` - Type: `string`.
* `poe_v` - Type: `bool`.
* `poecurr` - Type: `int`.
* `poepower` - Type: `int`.
* `poevolt` - Type: `int`.
* `power_class` - Type: `int`.
* `power_cycle_after` - Type: `string`.
* `power_cycle_host_alive` - Type: `bool`.
* `rate` - Type: `string`.
* `running` - Type: `bool`.
* `rx_align_error` - Type: `string`.
* `rx_broadcast` - Type: `int`.
* `rx_bytes` - Type: `int`.
* `rx_carrier_error` - Type: `string`.
* `rx_code_error` - Type: `string`.
* `rx_control` - Type: `string`.
* `rx_drop` - Type: `string`.
* `rx_error_events` - Type: `string`.
* `rx_fcs_error` - Type: `string`.
* `rx_fragment` - Type: `string`.
* `rx_jabber` - Type: `string`.
* `rx_length_error` - Type: `string`.
* `rx_loss` - Type: `bool`.
* `rx_multicast` - Type: `int`.
* `rx_overflow` - Type: `string`.
* `rx_packet` - Type: `int`.
* `rx_pause` - Type: `string`.
* `rx_power` - Type: `string`.
* `rx_too_long` - Type: `string`.
* `rx_too_short` - Type: `string`.
* `rx_unicast` - Type: `string`.
* `rx_unknown_op` - Type: `string`.
* `sfp_supported` - Type: `string`.
* `sfprate` - Type: `int`.
* `sfpshutdown` - Type: `bool`.
* `sm_link_length` - Type: `string`.
* `status` - Type: `enum(|off|on|disabled)`.
* `supply_voltage` - Type: `string`.
* `supported` - Type: `string`.
* `temperature` - Type: `string`.
* `tx_bias_current` - Type: `int`.
* `tx_broadcast` - Type: `int`.
* `tx_bytes` - Type: `int`.
* `tx_collision` - Type: `string`.
* `tx_control` - Type: `string`.
* `tx_deferred` - Type: `string`.
* `tx_drop` - Type: `string`.
* `tx_excessive_collision` - Type: `string`.
* `tx_excessive_deferred` - Type: `string`.
* `tx_fault` - Type: `bool`.
* `tx_fcs_error` - Type: `string`.
* `tx_fragment` - Type: `string`.
* `tx_jabber` - Type: `string`.
* `tx_late_collision` - Type: `string`.
* `tx_multicast` - Type: `int`.
* `tx_multiple_collision` - Type: `string`.
* `tx_packet` - Type: `int`.
* `tx_pause` - Type: `string`.
* `tx_pause_honorred` - Type: `string`.
* `tx_power` - Type: `string`.
* `tx_rx_1024_1518` - Type: `string`.
* `tx_rx_1024_max` - Type: `string`.
* `tx_rx_128_255` - Type: `string`.
* `tx_rx_1519_max` - Type: `string`.
* `tx_rx_256_511` - Type: `string`.
* `tx_rx_512_1023` - Type: `string`.
* `tx_rx_64` - Type: `string`.
* `tx_rx_65_127` - Type: `string`.
* `tx_rx_bytes` - Type: `string`.
* `tx_rx_packets` - Type: `string`.
* `tx_single_collision` - Type: `string`.
* `tx_too_short` - Type: `string`.
* `tx_total_collision` - Type: `string`.
* `tx_underrun` - Type: `string`.
* `tx_unicast` - Type: `string`.
* `vendor_name` - Type: `string`.
* `vendor_part_number` - Type: `string`.
* `vendor_revision` - Type: `string`.
* `vendor_serial` - Type: `string`.
* `wavelength` - Type: `string`.

