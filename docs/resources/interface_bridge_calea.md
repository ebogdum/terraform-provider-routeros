---
subcategory: "Bridge"
page_title: "RouterOS: routeros_interface_bridge_calea"
description: |-
  Bridge CALEA creates a session that drops the management connection on CHR. Skipped.
---

# Resource: routeros_interface_bridge_calea

Bridge CALEA creates a session that drops the management connection on CHR. Skipped.

## Example Usage

```terraform
resource "routeros_interface_bridge_calea" "calea_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `action` - (Optional) Type: `string`. RouterOS `action`.
* `arp_dst_address` - (Optional) Type: `string`. RouterOS `arp-dst-address`.
* `arp_dst_mac_address` - (Optional) Type: `string`. RouterOS `arp-dst-mac-address`.
* `arp_gratuitous` - (Optional) Type: `string`. RouterOS `arp-gratuitous`.
* `arp_hardware_type` - (Optional) Type: `string`. RouterOS `arp-hardware-type`.
* `arp_opcode` - (Optional) Type: `string`. RouterOS `arp-opcode`.
* `arp_packet_type` - (Optional) Type: `string`. RouterOS `arp-packet-type`.
* `arp_src_address` - (Optional) Type: `string`. RouterOS `arp-src-address`.
* `arp_src_mac_address` - (Optional) Type: `string`. RouterOS `arp-src-mac-address`.
* `chain` - (Optional) Type: `string`. RouterOS `chain`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dst_address` - (Optional) Type: `string`. RouterOS `dst-address`.
* `dst_address6` - (Optional) Type: `string`. RouterOS `dst-address6`.
* `dst_mac_address` - (Optional) Type: `string`. RouterOS `dst-mac-address`.
* `dst_port` - (Optional) Type: `string`. RouterOS `dst-port`.
* `in_bridge` - (Optional) Type: `string`. RouterOS `in-bridge`.
* `in_bridge_list` - (Optional) Type: `string`. RouterOS `in-bridge-list`.
* `in_interface` - (Optional) Type: `string`. RouterOS `in-interface`.
* `in_interface_list` - (Optional) Type: `string`. RouterOS `in-interface-list`.
* `ingress_priority` - (Optional) Type: `string`. RouterOS `ingress-priority`.
* `ip_protocol` - (Optional) Type: `string`. RouterOS `ip-protocol`.
* `limit` - (Optional) Type: `string`. RouterOS `limit`.
* `log` - (Optional) Type: `string`. RouterOS `log`.
* `log_prefix` - (Optional) Type: `string`. RouterOS `log-prefix`.
* `mac_protocol` - (Optional) Type: `string`. RouterOS `mac-protocol`.
* `out_bridge` - (Optional) Type: `string`. RouterOS `out-bridge`.
* `out_bridge_list` - (Optional) Type: `string`. RouterOS `out-bridge-list`.
* `out_interface` - (Optional) Type: `string`. RouterOS `out-interface`.
* `out_interface_list` - (Optional) Type: `string`. RouterOS `out-interface-list`.
* `packet_mark` - (Optional) Type: `string`. RouterOS `packet-mark`.
* `packet_type` - (Optional) Type: `string`. RouterOS `packet-type`.
* `sniff_id` - (Optional) Type: `string`. RouterOS `sniff-id`.
* `sniff_target` - (Optional) Type: `string`. RouterOS `sniff-target`.
* `sniff_target_port` - (Optional) Type: `string`. RouterOS `sniff-target-port`.
* `src_address` - (Optional) Type: `string`. RouterOS `src-address`.
* `src_address6` - (Optional) Type: `string`. RouterOS `src-address6`.
* `src_mac_address` - (Optional) Type: `string`. RouterOS `src-mac-address`.
* `src_port` - (Optional) Type: `string`. RouterOS `src-port`.
* `stp_flags` - (Optional) Type: `string`. RouterOS `stp-flags`.
* `stp_forward_delay` - (Optional) Type: `string`. RouterOS `stp-forward-delay`.
* `stp_hello_time` - (Optional) Type: `string`. RouterOS `stp-hello-time`.
* `stp_max_age` - (Optional) Type: `string`. RouterOS `stp-max-age`.
* `stp_msg_age` - (Optional) Type: `string`. RouterOS `stp-msg-age`.
* `stp_port` - (Optional) Type: `string`. RouterOS `stp-port`.
* `stp_root_address` - (Optional) Type: `string`. RouterOS `stp-root-address`.
* `stp_root_cost` - (Optional) Type: `string`. RouterOS `stp-root-cost`.
* `stp_root_priority` - (Optional) Type: `string`. RouterOS `stp-root-priority`.
* `stp_sender_address` - (Optional) Type: `string`. RouterOS `stp-sender-address`.
* `stp_sender_priority` - (Optional) Type: `string`. RouterOS `stp-sender-priority`.
* `stp_type` - (Optional) Type: `string`. RouterOS `stp-type`.
* `tls_host` - (Optional) Type: `string`. RouterOS `tls-host`.
* `vlan_encap` - (Optional) Type: `string`. RouterOS `vlan-encap`.
* `vlan_id` - (Optional) Type: `string`. RouterOS `vlan-id`.
* `vlan_priority` - (Optional) Type: `string`. RouterOS `vlan-priority`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_bridge_calea.example '*3'

# Named router
terraform import routeros_interface_bridge_calea.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_bridge_calea.example 'home/my-resource-name'
```
