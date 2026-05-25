---
subcategory: "System"
page_title: "RouterOS: routeros_system_resource_irq"
description: |-
  RouterOS resource.
---

# Data Source: routeros_system_resource_irq

Manages the RouterOS `/system/resource/irq` menu.

## Example Usage

```terraform
data "routeros_system_resource_irq" "irq_example" {
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
* `cpu` - (Optional) Type: `enum(auto)`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `active_cpu` - Type: `int`.
* `count` - Type: `int`.
* `irq` - Type: `int`.
* `per_cpu_count` - Type: `list`.
* `read_only` - Type: `bool`.
* `users` - Type: `string`.

