---
subcategory: "System & misc"
page_title: "RouterOS: routeros_terminal_cuu"
description: |-
  RouterOS resource.
---

# Resource: routeros_terminal_cuu

Manages the RouterOS `/terminal/cuu` menu.

## Example Usage

```terraform
resource "routeros_terminal_cuu" "cuu_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

