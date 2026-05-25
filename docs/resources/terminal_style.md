---
page_title: "RouterOS: routeros_terminal_style"
description: |-
  RouterOS resource.
---

# Resource: routeros_terminal_style

Manages the RouterOS `/terminal/style` menu.

## Example Usage

```terraform
resource "routeros_terminal_style" "style_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

