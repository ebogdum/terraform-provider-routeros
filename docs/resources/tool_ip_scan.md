---
page_title: "RouterOS: routeros_tool_ip_scan"
description: |-
  Long-running network scan; REST 60-second cap closes the session.
---

# Resource: routeros_tool_ip_scan

Long-running network scan; REST 60-second cap closes the session.

## Example Usage

```terraform
resource "routeros_tool_ip_scan" "ip_scan_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

