---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_traffic_generator_packet_template"
description: |-
  RouterOS resource.
---

# Resource: routeros_tool_traffic_generator_packet_template

Manages the RouterOS `/tool/traffic-generator/packet-template` menu.

## Example Usage

```terraform
resource "routeros_tool_traffic_generator_packet_template" "packet_template_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # data = "uninitialized"
  # data_byte = 0
  # dscp_ecn = "replace-me"
  # dst = "replace-me"
  # dst_port = "443"
  # flow_label = "replace-me"
  # frag_offset = "replace-me"
  # gateway = "replace-me"
  # header = "replace-me"
  # header_stack = "replace-me"
  # hop_limit = "replace-me"
  # interface = "ether1"
  # ip = "replace-me"
  # ip_id = "replace-me"
  # ipv6 = "replace-me"
  # mac = "replace-me"
  # name = "tf-example"
  # next_header = "replace-me"
  # port = "443"
  # priority = "replace-me"
  # protocol = "replace-me"
  # raw = "replace-me"
  # raw_packet_templates = "replace-me"
  # specbyte = "replace-me"
  # src = "replace-me"
  # src_port = "443"
  # tcp = "replace-me"
  # tcp_ack = "replace-me"
  # tcp_data_offset = "replace-me"
  # tcp_dst_port = "443"
  # tcp_flags = "replace-me"
  # tcp_src_port = "443"
  # tcp_syn = "replace-me"
  # tcp_urgent_pointer = "replace-me"
  # tcp_window_size = "replace-me"
  # traffic_class = "replace-me"
  # ttl = "replace-me"
  # udp = "replace-me"
  # vlan = "replace-me"
  # vlan_id = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `assumed_dscp_ecn` - (Read-only) Type: `string`.
* `assumed_dst` - (Read-only) Type: `string`.
* `assumed_dst_port` - (Read-only) Type: `string`.
* `assumed_flow_label` - (Read-only) Type: `string`.
* `assumed_frag_offset` - (Read-only) Type: `string`.
* `assumed_header` - (Read-only) Type: `string`.
* `assumed_interface` - (Read-only) Type: `string`.
* `assumed_ip_id` - (Read-only) Type: `string`.
* `assumed_next_header` - (Read-only) Type: `string`.
* `assumed_port` - (Read-only) Type: `string`.
* `assumed_priority` - (Read-only) Type: `string`.
* `assumed_protocol` - (Read-only) Type: `string`.
* `assumed_src` - (Read-only) Type: `string`.
* `assumed_src_port` - (Read-only) Type: `string`.
* `assumed_tcp_ack` - (Read-only) Type: `string`.
* `assumed_tcp_data_offset` - (Read-only) Type: `string`.
* `assumed_tcp_dst_port` - (Read-only) Type: `string`.
* `assumed_tcp_flags` - (Read-only) Type: `string`.
* `assumed_tcp_src_port` - (Read-only) Type: `string`.
* `assumed_tcp_syn` - (Read-only) Type: `string`.
* `assumed_tcp_urgent_pointer` - (Read-only) Type: `string`.
* `assumed_tcp_window_size` - (Read-only) Type: `string`.
* `assumed_traffic_class` - (Read-only) Type: `string`.
* `assumed_ttl` - (Read-only) Type: `string`.
* `assumed_vlan_id` - (Read-only) Type: `string`.
* `comment` - (Optional) Type: `string`.
* `compute_checksum_from_offset` - (Optional) Type: `string`. RouterOS `compute-checksum-from-offset`.
* `data` - (Optional) Type: `string`.
* `data_byte` - (Optional) Type: `int`.
* `dscp_ecn` - (Read-only) Type: `string`.
* `dst` - (Read-only) Type: `string`.
* `dst_port` - (Read-only) Type: `string`.
* `flow_label` - (Read-only) Type: `string`.
* `frag_offset` - (Read-only) Type: `string`.
* `gateway` - (Read-only) Type: `string`.
* `header` - (Read-only) Type: `string`.
* `header_stack` - (Optional) Type: `string`.
* `hop_limit` - (Read-only) Type: `string`.
* `interface` - (Optional) Type: `string`.
* `ip` - (Read-only) Type: `string`.
* `ip_dscp` - (Optional) Type: `string`. RouterOS `ip-dscp`.
* `ip_dst` - (Optional) Type: `string`. RouterOS `ip-dst`.
* `ip_frag_off` - (Optional) Type: `string`. RouterOS `ip-frag-off`.
* `ip_gateway` - (Optional) Type: `string`. RouterOS `ip-gateway`.
* `ip_id` - (Optional) Type: `string`.
* `ip_protocol` - (Optional) Type: `string`. RouterOS `ip-protocol`.
* `ip_src` - (Optional) Type: `string`. RouterOS `ip-src`.
* `ip_ttl` - (Optional) Type: `string`. RouterOS `ip-ttl`.
* `ipv6` - (Read-only) Type: `string`.
* `ipv6_dst` - (Optional) Type: `string`. RouterOS `ipv6-dst`.
* `ipv6_flow_label` - (Optional) Type: `string`. RouterOS `ipv6-flow-label`.
* `ipv6_gateway` - (Optional) Type: `string`. RouterOS `ipv6-gateway`.
* `ipv6_hop_limit` - (Optional) Type: `string`. RouterOS `ipv6-hop-limit`.
* `ipv6_next_header` - (Optional) Type: `string`. RouterOS `ipv6-next-header`.
* `ipv6_src` - (Optional) Type: `string`. RouterOS `ipv6-src`.
* `ipv6_traffic_class` - (Optional) Type: `string`. RouterOS `ipv6-traffic-class`.
* `mac` - (Read-only) Type: `string`.
* `mac_dst` - (Optional) Type: `string`. RouterOS `mac-dst`.
* `mac_protocol` - (Optional) Type: `string`. RouterOS `mac-protocol`.
* `mac_src` - (Optional) Type: `string`. RouterOS `mac-src`.
* `name` - (Optional) Type: `string`.
* `next_header` - (Read-only) Type: `string`.
* `port` - (Optional) Type: `string`.
* `priority` - (Read-only) Type: `string`.
* `protocol` - (Read-only) Type: `string`.
* `random_byte_offsets_and_masks` - (Optional) Type: `string`. RouterOS `random-byte-offsets-and-masks`.
* `random_ranges` - (Optional) Type: `string`. RouterOS `random-ranges`.
* `raw` - (Read-only) Type: `string`.
* `raw_header` - (Optional) Type: `string`. RouterOS `raw-header`.
* `raw_packet_templates` - (Read-only) Type: `string`.
* `specbyte` - (Read-only) Type: `string`.
* `special_footer` - (Optional) Type: `string`. RouterOS `special-footer`.
* `src` - (Read-only) Type: `string`.
* `src_port` - (Read-only) Type: `string`.
* `tcp` - (Read-only) Type: `string`.
* `tcp_ack` - (Optional) Type: `string`.
* `tcp_data_offset` - (Optional) Type: `string`.
* `tcp_dst_port` - (Optional) Type: `string`.
* `tcp_flags` - (Optional) Type: `string`.
* `tcp_src_port` - (Optional) Type: `string`.
* `tcp_syn` - (Optional) Type: `string`.
* `tcp_urgent_pointer` - (Optional) Type: `string`.
* `tcp_window_size` - (Optional) Type: `string`.
* `traffic_class` - (Read-only) Type: `string`.
* `ttl` - (Read-only) Type: `string`.
* `udp` - (Read-only) Type: `string`.
* `udp_checksum` - (Optional) Type: `string`. RouterOS `udp-checksum`.
* `udp_dst_port` - (Optional) Type: `string`. RouterOS `udp-dst-port`.
* `udp_src_port` - (Optional) Type: `string`. RouterOS `udp-src-port`.
* `vlan` - (Read-only) Type: `string`.
* `vlan_id` - (Optional) Type: `string`.
* `vlan_priority` - (Optional) Type: `string`. RouterOS `vlan-priority`.
* `vlan_protocol` - (Optional) Type: `string`. RouterOS `vlan-protocol`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_tool_traffic_generator_packet_template.example '*3'

# Named router
terraform import routeros_tool_traffic_generator_packet_template.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_tool_traffic_generator_packet_template.example 'home/my-resource-name'
```
