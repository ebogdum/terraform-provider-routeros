---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_list"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_list

Manages the RouterOS `/interface/list` menu.

## Example Usage

```terraform
resource "routeros_interface_list" "list_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # exclude = "replace-me"
  # include = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `builtin` - (Read-only) Type: `bool`.
* `comment` - (Optional) Type: `string`.
* `dynamic` - (Read-only) Type: `bool`.
* `exclude` - (Optional) Type: `string`. Defines interface list which members are excluded from the list. It is possible to add multiple lists separated by commas
* `include` - (Optional) Type: `string`. Defines interface list which members are included in the list. It is possible to add multiple lists separated by commas
* `name` - (Required) Type: `string`. Name of the interface list

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_list.example '*3'

# Named router
terraform import routeros_interface_list.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_list.example 'home/my-resource-name'
```
