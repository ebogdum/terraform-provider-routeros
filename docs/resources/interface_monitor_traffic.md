---
page_title: "RouterOS: routeros_interface_monitor_traffic"
description: |-
  Streaming monitor; the REST 60-second cap closes the session before output is produced.
---

# Resource: routeros_interface_monitor_traffic

Streaming monitor; the REST 60-second cap closes the session before output is produced.

## Example Usage

```terraform
resource "routeros_interface_monitor_traffic" "monitor_traffic_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

