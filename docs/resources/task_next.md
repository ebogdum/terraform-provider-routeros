---
page_title: "RouterOS: routeros_task_next"
description: |-
  Needs an active task. Skipped.
---

# Resource: routeros_task_next

Needs an active task. Skipped.

## Example Usage

```terraform
resource "routeros_task_next" "next_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

