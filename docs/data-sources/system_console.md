---
subcategory: "System"
page_title: "RouterOS: routeros_system_console"
description: |-
  Active console sessions -- RouterOS-managed; PUT EOFs because the endpoint isn't add-able.
---

# Data Source: routeros_system_console

Active console sessions -- RouterOS-managed; PUT EOFs because the endpoint isn't add-able.

## Example Usage

```terraform
data "routeros_system_console" "console_example" {
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
* `channel` - (Optional) Type: `int`.
* `disabled` - (Optional) Type: `bool`.
* `port` - (Optional) Type: `string`.
* `term` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `default` - Type: `bool`.
* `free` - Type: `bool`.
* `vcno` - Type: `int`.

