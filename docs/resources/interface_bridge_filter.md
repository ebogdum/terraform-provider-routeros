---
subcategory: "Bridge"
page_title: "RouterOS: routeros_interface_bridge_filter"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_bridge_filter

Manages the RouterOS `/interface/bridge/filter` menu.

## Example Usage

```terraform
resource "routeros_interface_bridge_filter" "filter_example" {
  # router = "my-router"  # which router to target; omit for the default
  chain = "forward"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "replace-me"
  # dst_address = "10.99.0.0/24"
  # dst_port = "443"
  # in_interface = "ether1"
  # in_interface_list = "LAN"
  # ingress_priority = "replace-me"
  # jump_target = "replace-me"
  # limit = "replace-me"
  # log = "replace-me"
  # log_prefix = "replace-me"
  # out_interface = "ether1"
  # out_interface_list = "LAN"
  # packet_mark = "replace-me"
  # src_address = "10.99.0.0/24"
  # src_mac_address = "10.99.0.0/24"
  # src_port = "443"
  # tls_host = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `action` - (Optional) Type: `string`.
* `arp_dst_address` - (Optional) Type: `string`. RouterOS `arp-dst-address`.
* `arp_dst_mac_address` - (Optional) Type: `string`. RouterOS `arp-dst-mac-address`.
* `arp_gratuitous` - (Optional) Type: `string`. RouterOS `arp-gratuitous`.
* `arp_hardware_type` - (Optional) Type: `string`. RouterOS `arp-hardware-type`.
* `arp_opcode` - (Optional) Type: `string`. RouterOS `arp-opcode`.
* `arp_packet_type` - (Optional) Type: `string`. RouterOS `arp-packet-type`.
* `arp_src_address` - (Optional) Type: `string`. RouterOS `arp-src-address`.
* `arp_src_mac_address` - (Optional) Type: `string`. RouterOS `arp-src-mac-address`.
* `chain` - (Required) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dst_address` - (Optional) Type: `string`.
* `dst_address6` - (Optional) Type: `string`. RouterOS `dst-address6`.
* `dst_mac_address` - (Optional) Type: `string`. RouterOS `dst-mac-address`.
* `dst_port` - (Optional) Type: `string`.
* `in_bridge` - (Optional) Type: `string`. RouterOS `in-bridge`.
* `in_bridge_list` - (Optional) Type: `string`. RouterOS `in-bridge-list`.
* `in_interface` - (Optional) Type: `string`.
* `in_interface_list` - (Optional) Type: `string`.
* `ingress_priority` - (Optional) Type: `string`.
* `ip_protocol` - (Optional) Type: `string`. RouterOS `ip-protocol`.
* `jump_target` - (Optional) Type: `string`.
* `limit` - (Optional) Type: `string`.
* `log` - (Optional) Type: `string`.
* `log_prefix` - (Optional) Type: `string`.
* `mac_protocol` - (Optional) Type: `string`. RouterOS `mac-protocol`.
* `new_packet_mark` - (Optional) Type: `string`. RouterOS `new-packet-mark`.
* `new_priority` - (Optional) Type: `string`. RouterOS `new-priority`.
* `out_bridge` - (Optional) Type: `string`. RouterOS `out-bridge`.
* `out_bridge_list` - (Optional) Type: `string`. RouterOS `out-bridge-list`.
* `out_interface` - (Optional) Type: `string`.
* `out_interface_list` - (Optional) Type: `string`.
* `packet_mark` - (Optional) Type: `string`.
* `packet_type` - (Optional) Type: `string`. RouterOS `packet-type`.
* `passthrough` - (Optional) Type: `string`. RouterOS `passthrough`.
* `src_address` - (Optional) Type: `string`.
* `src_address6` - (Optional) Type: `string`. RouterOS `src-address6`.
* `src_mac_address` - (Optional) Type: `string`.
* `src_port` - (Optional) Type: `string`.
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
* `tls_host` - (Optional) Type: `string`.
* `vlan_encap` - (Optional) Type: `string`. RouterOS `vlan-encap`.
* `vlan_id` - (Optional) Type: `string`. RouterOS `vlan-id`.
* `vlan_priority` - (Optional) Type: `string`. RouterOS `vlan-priority`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_bridge_filter.example '*3'

# Named router
terraform import routeros_interface_bridge_filter.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_bridge_filter.example 'home/my-resource-name'
```
