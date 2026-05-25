---
subcategory: "Terminal"
page_title: "RouterOS: routeros_terminal_el"
description: |-
  RouterOS resource.
---

# Resource: routeros_terminal_el

Manages the RouterOS `/terminal/el` menu.

## Example Usage

```terraform
resource "routeros_terminal_el" "el_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

