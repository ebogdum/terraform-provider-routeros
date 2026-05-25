---
page_title: "RouterOS: routeros_system_console"
description: |-
  Active console sessions -- RouterOS-managed; PUT EOFs because the endpoint isn't add-able.
---

# Resource: routeros_system_console

Active console sessions -- RouterOS-managed; PUT EOFs because the endpoint isn't add-able.

## Example Usage

```terraform
resource "routeros_system_console" "console_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # channel = 0
  # port = "443"
  # term = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `channel` - (Optional) Type: `int`.
* `disabled` - (Optional) Type: `bool`.
* `port` - (Optional) Type: `string`.
* `term` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `default` - Type: `bool`.
* `free` - Type: `bool`.
* `vcno` - Type: `int`.

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
