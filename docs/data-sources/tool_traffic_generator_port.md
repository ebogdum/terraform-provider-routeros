---
page_title: "RouterOS: routeros_tool_traffic_generator_port"
description: |-
  Discovered; needs tgen config
---

# Data Source: routeros_tool_traffic_generator_port

Discovered; needs tgen config

## Example Usage

```terraform
data "routeros_tool_traffic_generator_port" "port_example" {
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
* `disabled` - (Optional) Type: `bool`.
* `interface` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

