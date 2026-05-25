---
page_title: "RouterOS: routeros_system_serial_terminal"
description: |-
  Needs an active serial port. Skipped.
---

# Resource: routeros_system_serial_terminal

Needs an active serial port. Skipped.

## Example Usage

```terraform
resource "routeros_system_serial_terminal" "serial_terminal_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

