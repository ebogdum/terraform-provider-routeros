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
  # inactive = false
  # interface = "ether1"
  # mesh = "replace-me"
  # path_cost = 10
  # port_type = "auto"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `dr_address` - (Read-only) Type: `string`.
* `dynamic` - (Read-only) Type: `bool`.
* `hello_interval` - (Optional) Type: `int`.
* `inactive` - (Read-only) Type: `bool`.
* `interface` - (Optional) Type: `string`.
* `mesh` - (Optional) Type: `string`.
* `path_cost` - (Optional) Type: `int`.
* `port_type` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
