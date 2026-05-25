---
page_title: "RouterOS: routeros_tool_flood_ping"
description: |-
  Restricted by device-mode on most CHR/x86 installs. Skipped.
---

# Resource: routeros_tool_flood_ping

Restricted by device-mode on most CHR/x86 installs. Skipped.

## Example Usage

```terraform
resource "routeros_tool_flood_ping" "flood_ping_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

