---
page_title: "RouterOS: routeros_task_terminate"
description: |-
  Needs an active task .id. Skipped.
---

# Resource: routeros_task_terminate

Needs an active task .id. Skipped.

## Example Usage

```terraform
resource "routeros_task_terminate" "terminate_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

