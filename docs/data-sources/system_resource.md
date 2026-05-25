---
subcategory: "System"
page_title: "RouterOS: routeros_system_resource"
description: |-
  RouterOS resource.
---

# Data Source: routeros_system_resource

Manages the RouterOS `/system/resource` menu.

## Example Usage

```terraform
data "routeros_system_resource" "resource_example" {
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
* `architecture_name` - (Optional) Type: `string`.
* `board_name` - (Optional) Type: `string`.
* `build_time` - (Optional) Type: `string`.
* `cpu` - (Optional) Type: `string`.
* `cpu_count` - (Optional) Type: `int`.
* `cpu_frequency` - (Optional) Type: `int`.
* `cpu_load` - (Optional) Type: `int`.
* `free_hdd_space` - (Optional) Type: `int`.
* `free_memory` - (Optional) Type: `int`.
* `platform` - (Optional) Type: `string`.
* `total_hdd_space` - (Optional) Type: `int`.
* `total_memory` - (Optional) Type: `int`.
* `uptime` - (Optional) Type: `duration`.
* `version` - (Optional) Type: `string`.
* `write_sect_since_reboot` - (Optional) Type: `int`.
* `write_sect_total` - (Optional) Type: `int`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

