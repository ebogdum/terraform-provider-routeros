---
subcategory: "System & misc"
page_title: "RouterOS: routeros_caps_man_actual_interface_configuration"
description: |-
  CAPsMAN dynamic table; read-only.
---

# Resource: routeros_caps_man_actual_interface_configuration

CAPsMAN dynamic table; read-only.

## Example Usage

```terraform
resource "routeros_caps_man_actual_interface_configuration" "actual_interface_configuration_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_caps_man_actual_interface_configuration.example '*3'

# Named router
terraform import routeros_caps_man_actual_interface_configuration.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_caps_man_actual_interface_configuration.example 'home/my-resource-name'
```
