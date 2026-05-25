---
page_title: "RouterOS: routeros_system_shutdown"
description: |-
  Powers off the router. Verified working against a CHR VM. Skipped from the
---

# Resource: routeros_system_shutdown

Powers off the router. Verified working against a CHR VM. Skipped from the
general suite because it terminates the test target.


## Example Usage

```terraform
resource "routeros_system_shutdown" "shutdown_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

