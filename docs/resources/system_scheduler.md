---
subcategory: "System"
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
  name = "example"
  on_event = ":put \"tick\""

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # interval = "1h"
  # policy = "replace-me"
  # start_date = "replace-me"
  # start_time = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interval` - (Optional) Type: `duration`. Default: `1h`.
* `name` - (Required) Type: `string`. Default: `tf-acc-sched`.
* `on_event` - (Required) Type: `string`. Default: `:put "tick"`.
* `policy` - (Optional) Type: `string`.
* `start_date` - (Optional) Type: `string`.
* `start_time` - (Optional) Type: `string`.

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
