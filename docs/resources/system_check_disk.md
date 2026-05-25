---
page_title: "RouterOS: routeros_system_check_disk"
description: |-
  RouterOS resource.
---

# Resource: routeros_system_check_disk

Manages the RouterOS `/system/check-disk` menu.

## Example Usage

```terraform
resource "routeros_system_check_disk" "check_disk_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

