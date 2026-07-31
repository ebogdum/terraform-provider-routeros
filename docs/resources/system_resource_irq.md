---
subcategory: "System"
page_title: "RouterOS: routeros_system_resource_irq"
description: |-
  Mirrors RouterOS /system/resource/irq.
---

# Resource: routeros_system_resource_irq

Mirrors RouterOS `/system/resource/irq`.

## Example Usage

```terraform
resource "routeros_system_resource_irq" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # cpu = "replace-me"
  # irq = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `cpu` - (Optional) Type: `string`. RouterOS `cpu`.
* `irq` - (Read-only) Type: `string`. RouterOS `irq`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
terraform import routeros_system_resource_irq.example 'home::*3'
```
