---
page_title: "RouterOS: routeros_snmp_send_trap"
description: |-
  RouterOS resource.
---

# Resource: routeros_snmp_send_trap

Manages the RouterOS `/snmp/send-trap` menu.

## Example Usage

```terraform
resource "routeros_snmp_send_trap" "send_trap_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

