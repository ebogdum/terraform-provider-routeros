---
subcategory: "Storage"
page_title: "RouterOS: routeros_partition"
description: |-
  RouterOS resource.
---

# Data Source: routeros_partition

Manages the RouterOS `/partition` menu.

## Example Usage

```terraform
data "routeros_partition" "partition_example" {
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
* `activate` - (Optional) Type: `string`.
* `active` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`.
* `fallback_to` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `running` - (Optional) Type: `bool`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `size` - Type: `int`.
* `version` - Type: `string`.

