---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_traffic_generator_raw_packet_template"
description: |-
  RouterOS resource.
---

# Resource: routeros_tool_traffic_generator_raw_packet_template

Manages the RouterOS `/tool/traffic-generator/raw-packet-template` menu.

## Example Usage

```terraform
resource "routeros_tool_traffic_generator_raw_packet_template" "raw_packet_template_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # data = "uninitialized"
  # data_byte = 0
  # header = "replace-me"
  # ip_header_offset = "replace-me"
  # ipv6_header_offset = "replace-me"
  # name = "tf-example"
  # port = "443"
  # random = "replace-me"
  # random_byte_offsets_and_masks = "replace-me"
  # random_ranges = "replace-me"
  # specbyte = "replace-me"
  # special_footer = false
  # udp_header_offset = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`.
* `compute_checksum_from_offset` - (Optional) Type: `string`. RouterOS `compute-checksum-from-offset`.
* `data` - (Optional) Type: `string`.
* `data_byte` - (Optional) Type: `int`.
* `dynamic` - (Read-only) Type: `bool`.
* `header` - (Optional) Type: `string`.
* `header_length` - (Read-only) Type: `int`.
* `ip_header_offset` - (Optional) Type: `string`.
* `ipv6_header_offset` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `port` - (Optional) Type: `string`.
* `random` - (Read-only) Type: `string`.
* `random_byte_offsets_and_masks` - (Optional) Type: `string`.
* `random_ranges` - (Optional) Type: `string`.
* `specbyte` - (Read-only) Type: `string`.
* `special_footer` - (Optional) Type: `bool`.
* `tcp_header_offset` - (Optional) Type: `string`. RouterOS `tcp-header-offset`.
* `udp_compute_checksum` - (Optional) Type: `string`. RouterOS `udp-compute-checksum`.
* `udp_header_offset` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_tool_traffic_generator_raw_packet_template.example '*3'

# Named router
terraform import routeros_tool_traffic_generator_raw_packet_template.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_tool_traffic_generator_raw_packet_template.example 'home/my-resource-name'
```
