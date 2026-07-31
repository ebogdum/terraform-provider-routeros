---
subcategory: "System"
page_title: "RouterOS: routeros_system_console"
description: |-
  Active console sessions — RouterOS-managed; PUT EOFs because the endpoint isn't add-able.
---

# Resource: routeros_system_console

Active console sessions — RouterOS-managed; PUT EOFs because the endpoint isn't add-able.

## Example Usage

```terraform
resource "routeros_system_console" "console_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # channel = 0
  # port = "443"
  # term = "replace-me"
  # used = false
  # wedged = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `channel` - (Optional) Type: `int`.
* `default` - (Read-only) Type: `bool`.
* `disabled` - (Optional) Type: `bool`.
* `free` - (Read-only) Type: `bool`.
* `port` - (Optional) Type: `string`.
* `term` - (Optional) Type: `string`.
* `used` - (Read-only) Type: `bool`.
* `vc` - (Read-only) Type: `int`.
* `vcno` - (Read-only) Type: `int`.
* `wedged` - (Read-only) Type: `bool`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_system_console.example '*3'

# Named router
terraform import routeros_system_console.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_system_console.example 'home/my-resource-name'
```
