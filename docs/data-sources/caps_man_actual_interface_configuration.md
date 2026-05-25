---
subcategory: "System & misc"
page_title: "RouterOS: routeros_caps_man_actual_interface_configuration"
description: |-
  CAPsMAN dynamic table; read-only.
---

# Data Source: routeros_caps_man_actual_interface_configuration

CAPsMAN dynamic table; read-only.

## Example Usage

```terraform
data "routeros_caps_man_actual_interface_configuration" "actual_interface_configuration_example" {
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

