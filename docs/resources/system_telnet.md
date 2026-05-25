---
page_title: "RouterOS: routeros_system_telnet"
description: |-
  Interactive; would hang. Skipped.
---

# Resource: routeros_system_telnet

Interactive; would hang. Skipped.

## Example Usage

```terraform
resource "routeros_system_telnet" "telnet_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

