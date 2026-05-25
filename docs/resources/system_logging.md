---
subcategory: "System"
page_title: "RouterOS: routeros_system_logging"
description: |-
  RouterOS resource.
---

# Resource: routeros_system_logging

Manages the RouterOS `/system/logging` menu.

## Example Usage

```terraform
resource "routeros_system_logging" "logging_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "replace-me"
  # prefix = "replace-me"
  # regex = "replace-me"
  # topics = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `action` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `prefix` - (Optional) Type: `string`.
* `regex` - (Optional) Type: `string`.
* `topics` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `default` - Type: `bool`.
* `invalid` - Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_system_logging.example '*3'

# Named router
terraform import routeros_system_logging.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_system_logging.example 'home/my-resource-name'
```
