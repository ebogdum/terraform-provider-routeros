---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_reset_counters"
description: |-
  Needs an interface .id. Skipped -- generic acc test cannot target the right interface safely.
---

# Resource: routeros_interface_reset_counters

Needs an interface .id. Skipped -- generic acc test cannot target the right interface safely.

## Example Usage

```terraform
resource "routeros_interface_reset_counters" "reset_counters_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

