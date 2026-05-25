---
page_title: "RouterOS: routeros_queue_monitor"
description: |-
  Streaming monitor; REST 60-second cap closes the session.
---

# Resource: routeros_queue_monitor

Streaming monitor; REST 60-second cap closes the session.

## Example Usage

```terraform
resource "routeros_queue_monitor" "monitor_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

