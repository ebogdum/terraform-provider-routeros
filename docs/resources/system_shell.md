---
subcategory: "System"
page_title: "RouterOS: routeros_system_shell"
description: |-
  Interactive shell; not suitable for non-interactive acc tests.
---

# Resource: routeros_system_shell

Interactive shell; not suitable for non-interactive acc tests.

## Example Usage

```terraform
resource "routeros_system_shell" "shell_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

