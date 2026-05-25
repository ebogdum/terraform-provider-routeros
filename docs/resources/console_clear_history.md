---
subcategory: "System & misc"
page_title: "RouterOS: routeros_console_clear_history"
description: |-
  RouterOS resource.
---

# Resource: routeros_console_clear_history

Manages the RouterOS `/console/clear-history` menu.

## Example Usage

```terraform
resource "routeros_console_clear_history" "clear_history_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

