---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_discourse"
description: |-
  Requires explicit dsts (destinations) list.
---

# Resource: routeros_routing_discourse

Requires explicit dsts (destinations) list.

## Example Usage

```terraform
resource "routeros_routing_discourse" "discourse_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

