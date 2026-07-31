---
subcategory: "Scripts & scheduler"
page_title: "RouterOS: routeros_system_script"
description: |-
  RouterOS resource.
---

# Resource: routeros_system_script

Manages the RouterOS `/system/script` menu.

## Example Usage

```terraform
resource "routeros_system_script" "script_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"
  source = ":put \"hello\""

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # dont_require_permissions = false
  # policy = []
  # run_script = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dont_require_permissions` - (Optional) Type: `bool`.
* `invalid` - (Read-only) Type: `bool`.
* `last_time_started` - (Read-only) Type: `string`.
* `name` - (Required) Type: `string`.
* `owner` - (Read-only) Type: `string`.
* `policy` - (Optional) Type: `list`.
* `run_count` - (Read-only) Type: `int`.
* `run_script` - (Read-only) Type: `string`.
* `source` - (Required) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_system_script.example '*3'

# Named router
terraform import routeros_system_script.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_system_script.example 'home/my-resource-name'
```
