---
subcategory: "Disks"
page_title: "RouterOS: routeros_disk_scan"
description: |-
  Needs a real disk.
---

# Resource: routeros_disk_scan

Needs a real disk.

## Example Usage

```terraform
resource "routeros_disk_scan" "scan_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

