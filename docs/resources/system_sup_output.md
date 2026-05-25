---
subcategory: "System"
page_title: "RouterOS: routeros_system_sup_output"
description: |-
  RouterOS resource.
---

# Resource: routeros_system_sup_output

Manages the RouterOS `/system/sup-output` menu.

## Example Usage

```terraform
resource "routeros_system_sup_output" "sup_output_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

