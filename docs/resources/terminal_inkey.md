---
page_title: "RouterOS: routeros_terminal_inkey"
description: |-
  Reads a keystroke from interactive terminal. Skipped.
---

# Resource: routeros_terminal_inkey

Reads a keystroke from interactive terminal. Skipped.

## Example Usage

```terraform
resource "routeros_terminal_inkey" "inkey_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

