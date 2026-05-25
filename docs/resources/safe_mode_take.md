---
page_title: "RouterOS: routeros_safe_mode_take"
description: |-
  RouterOS resource.
---

# Resource: routeros_safe_mode_take

Manages the RouterOS `/safe-mode/take` menu.

## Example Usage

```terraform
resource "routeros_safe_mode_take" "take_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

