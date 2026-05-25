---
page_title: "RouterOS: routeros_certificate_unset"
description: |-
  Needs cert .id + field. Skipped.
---

# Resource: routeros_certificate_unset

Needs cert .id + field. Skipped.

## Example Usage

```terraform
resource "routeros_certificate_unset" "unset_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

