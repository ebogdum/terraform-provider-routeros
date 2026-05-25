---
page_title: "RouterOS: routeros_tool_romon_port"
description: |-
  RoMON port references a specific interface; values vary per device. Skipped.
---

# Data Source: routeros_tool_romon_port

RoMON port references a specific interface; values vary per device. Skipped.

## Example Usage

```terraform
data "routeros_tool_romon_port" "port_example" {
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
* `cost` - (Optional) Type: `int`. Default: `100`.
* `disabled` - (Optional) Type: `bool`.
* `forbid` - (Optional) Type: `bool`.
* `interface` - (Optional) Type: `string`.
* `secrets` - (Optional) Type: `string`. **Sensitive.**

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `default` - Type: `bool`.
* `dynamic` - Type: `bool`.

