---
page_title: "RouterOS: routeros_tool_netwatch"
description: |-
  RouterOS resource.
---

# Data Source: routeros_tool_netwatch

Manages the RouterOS `/tool/netwatch` menu.

## Example Usage

```terraform
data "routeros_tool_netwatch" "netwatch_example" {
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
* `certificate` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dns_server` - (Optional) Type: `string`.
* `host` - (Required) Type: `string`. Default: `127.0.0.1`.
* `interval` - (Optional) Type: `duration`. Default: `1m`.
* `name` - (Optional) Type: `string`.
* `port` - (Optional) Type: `string`.
* `src_address` - (Optional) Type: `string`.
* `timeout` - (Optional) Type: `string`.
* `ttl` - (Optional) Type: `string`.
* `type` - (Optional) Type: `string`. Default: `icmp`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

