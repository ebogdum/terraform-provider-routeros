---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_mesh_port"
description: |-
  Discovered via WebFig; needs mesh-interface fixture
---

# Data Source: routeros_interface_mesh_port

Discovered via WebFig; needs mesh-interface fixture

## Example Usage

```terraform
data "routeros_interface_mesh_port" "port_example" {
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
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `hello_interval` - (Optional) Type: `int`. Default: `10`.
* `interface` - (Optional) Type: `string`.
* `mesh` - (Optional) Type: `string`.
* `path_cost` - (Optional) Type: `int`. Default: `10`.
* `port_type` - (Optional) Type: `enum(auto|WDS|wireless|ethernet)`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

