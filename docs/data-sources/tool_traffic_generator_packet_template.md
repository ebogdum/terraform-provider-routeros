---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_traffic_generator_packet_template"
description: |-
  RouterOS resource.
---

# Data Source: routeros_tool_traffic_generator_packet_template

Manages the RouterOS `/tool/traffic-generator/packet-template` menu.

## Example Usage

```terraform
data "routeros_tool_traffic_generator_packet_template" "packet_template_example" {
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

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
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

