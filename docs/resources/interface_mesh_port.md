---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_mesh_port"
description: |-
  Discovered via WebFig; needs mesh-interface fixture
---

# Resource: routeros_interface_mesh_port

Discovered via WebFig; needs mesh-interface fixture

## Example Usage

```terraform
resource "routeros_interface_mesh_port" "port_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # hello_interval = 10
  # interface = "ether1"
  # mesh = "replace-me"
  # path_cost = 10
  # port_type = "auto"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `hello_interval` - (Optional) Type: `int`. Default: `10`.
* `interface` - (Optional) Type: `string`.
* `mesh` - (Optional) Type: `string`.
* `path_cost` - (Optional) Type: `int`. Default: `10`.
* `port_type` - (Optional) Type: `enum(auto|WDS|wireless|ethernet)`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_mesh_port.example '*3'

# Named router
terraform import routeros_interface_mesh_port.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_mesh_port.example 'home/my-resource-name'
```
