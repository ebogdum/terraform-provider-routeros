---
subcategory: "Scripts & scheduler"
page_title: "RouterOS: routeros_system_scheduler"
description: |-
  RouterOS resource.
---

# Resource: routeros_system_scheduler

Manages the RouterOS `/system/scheduler` menu.

## Example Usage

```terraform
resource "routeros_system_scheduler" "scheduler_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"
  on_event = ":put \"tick\""

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # interval = "1h"
  # policy = "replace-me"
  # start_date = "replace-me"
  # start_time = "23:57:05"  # or "startup" to run once at boot
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interval` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `next_run` - (Read-only) Type: `string`.
* `on_event` - (Required) Type: `string`.
* `owner` - (Read-only) Type: `string`.
* `policy` - (Optional) Type: `string`.
* `run_count` - (Read-only) Type: `int`.
* `start_date` - (Optional) Type: `string`.
* `start_time` - (Optional) Type: `string`. A specific `HH:MM:SS` time, or the keyword `startup` to run once when RouterOS boots. RouterOS also accepts shorter spellings such as `23:57` and `0:0:0`, but rewrites them, so they are refused here rather than applied as something else.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_system_scheduler.example '*3'

# Named router
terraform import routeros_system_scheduler.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_system_scheduler.example 'home/my-resource-name'
```
