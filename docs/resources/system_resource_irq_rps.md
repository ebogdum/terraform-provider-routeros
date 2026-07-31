---
subcategory: "System"
page_title: "RouterOS: routeros_system_resource_irq_rps"
description: |-
  Mirrors RouterOS /system/resource/irq/rps.
---

# Resource: routeros_system_resource_irq_rps

Mirrors RouterOS `/system/resource/irq/rps`.

## Example Usage

```terraform
resource "routeros_system_resource_irq_rps" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # disabled = true
  # name = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `disabled` - (Optional) Type: `bool`. RouterOS `disabled`.
* `name` - (Read-only) Type: `string`. RouterOS `name`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
terraform import routeros_system_resource_irq_rps.example 'home::*3'
```
