---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_traffic_generator_raw_packet_template"
description: |-
  RouterOS resource.
---

# Data Source: routeros_tool_traffic_generator_raw_packet_template

Manages the RouterOS `/tool/traffic-generator/raw-packet-template` menu.

## Example Usage

```terraform
data "routeros_tool_traffic_generator_raw_packet_template" "raw_packet_template_example" {
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
* `data` - (Optional) Type: `enum(uninitialized|random|specific byte|incrementing)`.
* `data_byte` - (Optional) Type: `int`.
* `header` - (Optional) Type: `string`.
* `ip_header_offset` - (Optional) Type: `string`.
* `ipv6_header_offset` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `port` - (Optional) Type: `string`.
* `random_byte_offsets_and_masks` - (Optional) Type: `string`.
* `random_ranges` - (Optional) Type: `string`.
* `special_footer` - (Optional) Type: `bool`.
* `udp_header_offset` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

