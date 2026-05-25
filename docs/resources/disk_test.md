---
page_title: "RouterOS: routeros_disk_test"
description: |-
  Needs a specific disk; not portable.
---

# Resource: routeros_disk_test

Needs a specific disk; not portable.

## Example Usage

```terraform
resource "routeros_disk_test" "test_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

