---
subcategory: "System & misc"
page_title: "RouterOS: routeros_console_inspect"
description: |-
  RouterOS resource.
---

# Resource: routeros_console_inspect

Manages the RouterOS `/console/inspect` menu.

## Example Usage

```terraform
resource "routeros_console_inspect" "inspect_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

