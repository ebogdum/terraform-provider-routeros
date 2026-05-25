---
subcategory: "RIP"
page_title: "RouterOS: routeros_routing_rip_instance"
description: |-
  RouterOS resource.
---

# Data Source: routeros_routing_rip_instance

Manages the RouterOS `/routing/rip/instance` menu.

## Example Usage

```terraform
data "routeros_routing_rip_instance" "instance_example" {
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
* `afi` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `input_filter` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `originate_default` - (Optional) Type: `string`.
* `output_filter` - (Optional) Type: `string`.
* `redistribute` - (Optional) Type: `string`.
* `route_gc_timeout` - (Optional) Type: `string`.
* `route_timeout` - (Optional) Type: `string`.
* `routing_table` - (Optional) Type: `string`.
* `select_output_filter` - (Optional) Type: `string`.
* `update_interval` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

