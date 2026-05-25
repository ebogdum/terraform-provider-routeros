---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_sniffer_packet"
description: |-
  RouterOS resource.
---

# Data Source: routeros_tool_sniffer_packet

Manages the RouterOS `/tool/sniffer/packet` menu.

## Example Usage

```terraform
data "routeros_tool_sniffer_packet" "packet_example" {
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
* `comment` - (Optional) Type: `string`. Free-form comment.
* `cpu` - (Optional) Type: `int`.
* `direction` - (Optional) Type: `enum(rx|tx)`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dst_address` - (Optional) Type: `string`.
* `dst_mac_address` - (Optional) Type: `string`.
* `dst_port` - (Optional) Type: `int`.
* `fragment_offset` - (Optional) Type: `int`.
* `identification` - (Optional) Type: `int`.
* `interface` - (Optional) Type: `string`.
* `ip_header_size` - (Optional) Type: `int`.
* `ip_packet_size` - (Optional) Type: `int`.
* `ip_protocol` - (Optional) Type: `int`.
* `num` - (Optional) Type: `int`.
* `protocol` - (Optional) Type: `int`.
* `raw_data` - (Optional) Type: `string`.
* `size` - (Optional) Type: `int`.
* `src_address` - (Optional) Type: `string`.
* `src_mac_address` - (Optional) Type: `string`.
* `src_port` - (Optional) Type: `int`.
* `time` - (Optional) Type: `string`.
* `tos` - (Optional) Type: `int`.
* `ttl` - (Optional) Type: `int`.
* `vlan_id` - (Optional) Type: `int`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

