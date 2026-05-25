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
* `backup` - (Optional) Type: `string`.
* `basename` - (Optional) Type: `string`.
* `container` - (Optional) Type: `int`. Default: `4.294967295e+09`.
* `contents` - (Optional) Type: `string`.
* `directory` - (Optional) Type: `string`.
* `file_name` - (Optional) Type: `string`.
* `hasvpn` - (Optional) Type: `string`.
* `invalid` - (Optional) Type: `string`.
* `invalidfile` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `nondir` - (Optional) Type: `string`.
* `restore` - (Optional) Type: `string`.
* `share` - (Optional) Type: `string`.
* `shared` - (Optional) Type: `bool`.
* `type` - (Optional) Type: `int`.
* `unshare` - (Optional) Type: `string`.
* `valid` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `family` - Type: `int`.
* `file_share_url` - Type: `string`.
* `last_modified` - Type: `string`.
* `size` - Type: `string`.

