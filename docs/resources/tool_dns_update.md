---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_dns_update"
description: |-
  Sends an RFC 2136 DDNS update; needs a real DDNS server.
---

# Resource: routeros_tool_dns_update

Sends an RFC 2136 DDNS update; needs a real DDNS server.

## Example Usage

```terraform
resource "routeros_tool_dns_update" "dns_update_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

