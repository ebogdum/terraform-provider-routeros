---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_sniffer_protocol"
description: |-
  RouterOS resource.
---

# Data Source: routeros_tool_sniffer_protocol

Manages the RouterOS `/tool/sniffer/protocol` menu.

## Example Usage

```terraform
data "routeros_tool_sniffer_protocol" "protocol_example" {
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
* `bytes` - (Optional) Type: `int`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `ip_protocol` - (Optional) Type: `int`.
* `packets` - (Optional) Type: `int`.
* `port` - (Optional) Type: `int`.
* `protocol` - (Optional) Type: `int`.
* `share` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

