---
page_title: "RouterOS: routeros_tool_ping_speed"
description: |-
  Long-running; REST 60-second cap closes the session.
---

# Resource: routeros_tool_ping_speed

Long-running; REST 60-second cap closes the session.

## Example Usage

```terraform
resource "routeros_tool_ping_speed" "ping_speed_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

