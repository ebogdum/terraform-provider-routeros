---
subcategory: "Files"
page_title: "RouterOS: routeros_file"
description: |-
  Creating a file via REST requires writing the contents in a follow-up call; not in the acc-test fast path.
---

# Data Source: routeros_file

Creating a file via REST requires writing the contents in a follow-up call; not in the acc-test fast path.

## Example Usage

```terraform
data "routeros_file" "file_example" {
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
* `contents` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `type` - (Optional) Type: `int`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `last_modified` - Type: `string`.

