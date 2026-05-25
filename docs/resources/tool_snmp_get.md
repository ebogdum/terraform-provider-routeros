---
page_title: "RouterOS: routeros_tool_snmp_get"
description: |-
  Needs a real SNMP agent. Skipped.
---

# Resource: routeros_tool_snmp_get

Needs a real SNMP agent. Skipped.

## Example Usage

```terraform
resource "routeros_tool_snmp_get" "snmp_get_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

