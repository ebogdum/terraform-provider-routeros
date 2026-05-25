---
subcategory: "Files"
page_title: "RouterOS: routeros_file_read"
description: |-
  Needs a specific file.
---

# Resource: routeros_file_read

Needs a specific file.

## Example Usage

```terraform
resource "routeros_file_read" "read_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

