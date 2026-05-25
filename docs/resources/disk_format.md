---
page_title: "RouterOS: routeros_disk_format"
description: |-
  Formats a block device. Requires the .id of a real /disk entry, which an
---

# Resource: routeros_disk_format

Formats a block device. Requires the .id of a real /disk entry, which an
automated test cannot create. Skipped.


## Example Usage

```terraform
resource "routeros_disk_format" "format_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

