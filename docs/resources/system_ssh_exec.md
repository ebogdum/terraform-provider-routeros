---
subcategory: "System"
page_title: "RouterOS: routeros_system_ssh_exec"
description: |-
  Needs a reachable SSH peer with a real command. Skipped.
---

# Resource: routeros_system_ssh_exec

Needs a reachable SSH peer with a real command. Skipped.

## Example Usage

```terraform
resource "routeros_system_ssh_exec" "ssh_exec_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

