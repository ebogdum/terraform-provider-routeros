---
subcategory: "Ping"
page_title: "RouterOS: routeros_ping"
description: |-
  Long-running by default; REST 60-second cap closes the session.
---

# Resource: routeros_ping

Long-running by default; REST 60-second cap closes the session.

## Example Usage

```terraform
resource "routeros_ping" "ping_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

