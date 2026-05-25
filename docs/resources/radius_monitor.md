---
subcategory: "RADIUS"
page_title: "RouterOS: routeros_radius_monitor"
description: |-
  Needs radius server .id. Skipped.
---

# Resource: routeros_radius_monitor

Needs radius server .id. Skipped.

## Example Usage

```terraform
resource "routeros_radius_monitor" "monitor_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

