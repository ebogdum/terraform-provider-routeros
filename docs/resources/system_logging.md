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
  # name = "tf-example"
  # prefix = "replace-me"
  # regex = "replace-me"
  # topics = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `action` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default` - (Read-only) Type: `bool`.
* `disabled` - (Optional) Type: `bool`.
* `invalid` - (Read-only) Type: `bool`.
* `name` - (Read-only) Type: `string`.
* `prefix` - (Optional) Type: `string`.
* `regex` - (Optional) Type: `string`.
* `topics` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
