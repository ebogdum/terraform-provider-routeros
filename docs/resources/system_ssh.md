---
subcategory: "System"
page_title: "RouterOS: routeros_system_ssh"
description: |-
  SSH client opens an interactive session; not suitable for non-interactive acc tests.
---

# Resource: routeros_system_ssh

SSH client opens an interactive session; not suitable for non-interactive acc tests.

## Example Usage

```terraform
resource "routeros_system_ssh" "ssh_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

