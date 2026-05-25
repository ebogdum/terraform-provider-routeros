---
page_title: "RouterOS: routeros_ip_traffic_flow_target"
description: |-
  Discovered; required dst-address must be valid
---

# Data Source: routeros_ip_traffic_flow_target

Discovered; required dst-address must be valid

## Example Usage

```terraform
data "routeros_ip_traffic_flow_target" "target_example" {
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
* `dst_address` - (Optional) Type: `string`.
* `port` - (Optional) Type: `int`. Default: `1234`.
* `src_address` - (Optional) Type: `string`.
* `version` - (Optional) Type: `enum(1|5|9|IPFIX)`. Default: `9`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

