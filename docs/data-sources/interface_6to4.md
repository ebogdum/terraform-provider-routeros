---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_6to4"
description: |-
  6to4 tunnel deletion races on CHR (DELETE returns errors even after success). Skipped from automated acc tests.
---

# Data Source: routeros_interface_6to4

6to4 tunnel deletion races on CHR (DELETE returns errors even after success). Skipped from automated acc tests.

## Example Usage

```terraform
data "routeros_interface_6to4" "6to4_example" {
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
* `clamp_tcp_mss` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dont_fragment` - (Optional) Type: `bool`.
* `dscp` - (Optional) Type: `string`.
* `local_address` - (Optional) Type: `ip`.
* `mtu` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `remote_address` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

