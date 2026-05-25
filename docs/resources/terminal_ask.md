---
page_title: "RouterOS: routeros_terminal_ask"
description: |-
  RouterOS resource.
---

# Resource: routeros_terminal_ask

Manages the RouterOS `/terminal/ask` menu.

## Example Usage

```terraform
resource "routeros_terminal_ask" "ask_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

