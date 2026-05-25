---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_mac_telnet"
description: |-
  Interactive MAC-Telnet session. Skipped.
---

# Resource: routeros_tool_mac_telnet

Interactive MAC-Telnet session. Skipped.

## Example Usage

```terraform
resource "routeros_tool_mac_telnet" "mac_telnet_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

