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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`.
* `data` - (Optional) Type: `enum(uninitialized|random|specific-byte|incrementing)`.
* `data_byte` - (Optional) Type: `int`.
* `header` - (Optional) Type: `string`.
* `ip_header_offset` - (Optional) Type: `string`.
* `ipv6_header_offset` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `port` - (Optional) Type: `string`.
* `random` - (Optional) Type: `string`.
* `random_byte_offsets_and_masks` - (Optional) Type: `string`.
* `random_ranges` - (Optional) Type: `string`.
* `specbyte` - (Optional) Type: `string`.
* `special_footer` - (Optional) Type: `bool`.
* `udp_header_offset` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `dynamic` - Type: `bool`.
* `header_length` - Type: `int`.

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
