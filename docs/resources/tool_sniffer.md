---
page_title: "RouterOS: routeros_tool_sniffer"
description: |-
  RouterOS resource.
---

# Resource: routeros_tool_sniffer

Manages the RouterOS `/tool/sniffer` menu.

## Example Usage

```terraform
resource "routeros_tool_sniffer" "sniffer_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # file_limit = 0
  # file_name = "replace-me"
  # filter_cpu = "replace-me"
  # filter_direction = "replace-me"
  # filter_dst_ip_address = "10.99.0.0/24"
  # filter_dst_ipv6_address = "10.99.0.0/24"
  # filter_dst_mac_address = "10.99.0.0/24"
  # filter_dst_port = "443"
  # filter_interface = "replace-me"
  # filter_ip_address = "10.99.0.0/24"
  # filter_ip_protocol = "replace-me"
  # filter_ipv6_address = "10.99.0.0/24"
  # filter_mac_address = "10.99.0.0/24"
  # filter_mac_protocol = "replace-me"
  # filter_operator_between_entries = "replace-me"
  # filter_port = "443"
  # filter_size = "replace-me"
  # filter_src_ip_address = "10.99.0.0/24"
  # filter_src_ipv6_address = "10.99.0.0/24"
  # filter_src_mac_address = "10.99.0.0/24"
  # filter_src_port = "443"
  # filter_stream = false
  # filter_vlan = "replace-me"
  # max_packet_size = 0
  # memory_limit = 0
  # memory_scroll = false
  # only_headers = false
  # quick_rows = 0
  # quick_show_frame = false
  # streaming_enabled = false
  # streaming_server = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `file_limit` - (Optional) Type: `int`.
* `file_name` - (Optional) Type: `string`.
* `filter_cpu` - (Optional) Type: `string`.
* `filter_direction` - (Optional) Type: `string`.
* `filter_dst_ip_address` - (Optional) Type: `string`.
* `filter_dst_ipv6_address` - (Optional) Type: `string`.
* `filter_dst_mac_address` - (Optional) Type: `string`.
* `filter_dst_port` - (Optional) Type: `string`.
* `filter_interface` - (Optional) Type: `string`.
* `filter_ip_address` - (Optional) Type: `string`.
* `filter_ip_protocol` - (Optional) Type: `string`.
* `filter_ipv6_address` - (Optional) Type: `string`.
* `filter_mac_address` - (Optional) Type: `string`.
* `filter_mac_protocol` - (Optional) Type: `string`.
* `filter_operator_between_entries` - (Optional) Type: `string`.
* `filter_port` - (Optional) Type: `string`.
* `filter_size` - (Optional) Type: `string`.
* `filter_src_ip_address` - (Optional) Type: `string`.
* `filter_src_ipv6_address` - (Optional) Type: `string`.
* `filter_src_mac_address` - (Optional) Type: `string`.
* `filter_src_port` - (Optional) Type: `string`.
* `filter_stream` - (Optional) Type: `bool`.
* `filter_vlan` - (Optional) Type: `string`.
* `max_packet_size` - (Optional) Type: `int`.
* `memory_limit` - (Optional) Type: `int`.
* `memory_scroll` - (Optional) Type: `bool`.
* `only_headers` - (Optional) Type: `bool`.
* `quick_rows` - (Optional) Type: `int`.
* `quick_show_frame` - (Optional) Type: `bool`.
* `streaming_enabled` - (Optional) Type: `bool`.
* `streaming_server` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `running` - Type: `bool`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_tool_sniffer.this 'home'
```
