---
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
  # header_stack = "replace-me"
  # interface = "ether1"
  # ip_id = "replace-me"
  # name = "tf-example"
  # port = "443"
  # tcp_ack = "replace-me"
  # tcp_data_offset = "replace-me"
  # tcp_dst_port = "443"
  # tcp_flags = "replace-me"
  # tcp_src_port = "443"
  # tcp_syn = "replace-me"
  # tcp_urgent_pointer = "replace-me"
  # tcp_window_size = "replace-me"
  # vlan_id = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`.
* `data` - (Optional) Type: `enum(uninitialized|random|specific byte|incrementing)`.
* `data_byte` - (Optional) Type: `int`.
* `header_stack` - (Optional) Type: `string`.
* `interface` - (Optional) Type: `string`.
* `ip_id` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `port` - (Optional) Type: `string`.
* `tcp_ack` - (Optional) Type: `string`.
* `tcp_data_offset` - (Optional) Type: `string`.
* `tcp_dst_port` - (Optional) Type: `string`.
* `tcp_flags` - (Optional) Type: `string`.
* `tcp_src_port` - (Optional) Type: `string`.
* `tcp_syn` - (Optional) Type: `string`.
* `tcp_urgent_pointer` - (Optional) Type: `string`.
* `tcp_window_size` - (Optional) Type: `string`.
* `vlan_id` - (Optional) Type: `string`.

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
