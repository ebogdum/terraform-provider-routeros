---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_lte"
description: |-
  LTE interfaces are physical-device backed; skipped on virtual devices.
---

# Data Source: routeros_interface_lte

LTE interfaces are physical-device backed; skipped on virtual devices.

## Example Usage

```terraform
data "routeros_interface_lte" "lte_example" {
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
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

