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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`.
* `data` - (Optional) Type: `enum(uninitialized|random|specific-byte|incrementing)`.
* `data_byte` - (Optional) Type: `int`.
* `dscp_ecn` - (Optional) Type: `string`.
* `dst` - (Optional) Type: `string`.
* `dst_port` - (Optional) Type: `string`.
* `flow_label` - (Optional) Type: `string`.
* `frag_offset` - (Optional) Type: `string`.
* `gateway` - (Optional) Type: `string`.
* `header` - (Optional) Type: `string`.
* `header_stack` - (Optional) Type: `string`.
* `hop_limit` - (Optional) Type: `string`.
* `interface` - (Optional) Type: `string`.
* `ip` - (Optional) Type: `string`.
* `ip_id` - (Optional) Type: `string`.
* `ipv6` - (Optional) Type: `string`.
* `mac` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `next_header` - (Optional) Type: `string`.
* `port` - (Optional) Type: `string`.
* `priority` - (Optional) Type: `string`.
* `protocol` - (Optional) Type: `string`.
* `raw` - (Optional) Type: `string`.
* `raw_packet_templates` - (Optional) Type: `string`.
* `specbyte` - (Optional) Type: `string`.
* `src` - (Optional) Type: `string`.
* `src_port` - (Optional) Type: `string`.
* `tcp` - (Optional) Type: `string`.
* `tcp_ack` - (Optional) Type: `string`.
* `tcp_data_offset` - (Optional) Type: `string`.
* `tcp_dst_port` - (Optional) Type: `string`.
* `tcp_flags` - (Optional) Type: `string`.
* `tcp_src_port` - (Optional) Type: `string`.
* `tcp_syn` - (Optional) Type: `string`.
* `tcp_urgent_pointer` - (Optional) Type: `string`.
* `tcp_window_size` - (Optional) Type: `string`.
* `traffic_class` - (Optional) Type: `string`.
* `ttl` - (Optional) Type: `string`.
* `udp` - (Optional) Type: `string`.
* `vlan` - (Optional) Type: `string`.
* `vlan_id` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `assumed_dscp_ecn` - Type: `string`.
* `assumed_dst` - Type: `string`.
* `assumed_dst_port` - Type: `string`.
* `assumed_flow_label` - Type: `string`.
* `assumed_frag_offset` - Type: `string`.
* `assumed_header` - Type: `string`.
* `assumed_interface` - Type: `string`.
* `assumed_ip_id` - Type: `string`.
* `assumed_next_header` - Type: `string`.
* `assumed_port` - Type: `string`.
* `assumed_priority` - Type: `string`.
* `assumed_protocol` - Type: `string`.
* `assumed_src` - Type: `string`.
* `assumed_src_port` - Type: `string`.
* `assumed_tcp_ack` - Type: `string`.
* `assumed_tcp_data_offset` - Type: `string`.
* `assumed_tcp_dst_port` - Type: `string`.
* `assumed_tcp_flags` - Type: `string`.
* `assumed_tcp_src_port` - Type: `string`.
* `assumed_tcp_syn` - Type: `string`.
* `assumed_tcp_urgent_pointer` - Type: `string`.
* `assumed_tcp_window_size` - Type: `string`.
* `assumed_traffic_class` - Type: `string`.
* `assumed_ttl` - Type: `string`.
* `assumed_vlan_id` - Type: `string`.

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
