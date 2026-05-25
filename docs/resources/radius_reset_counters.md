---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_radius_reset_counters"
description: |-
  RouterOS resource.
---

# Resource: routeros_radius_reset_counters

Manages the RouterOS `/radius/reset-counters` menu.

## Example Usage

```terraform
resource "routeros_radius_reset_counters" "reset_counters_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

