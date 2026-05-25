---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_bandwidth_test"
description: |-
  Needs a remote bandwidth-server. Skipped (long-running, network-dependent).
---

# Resource: routeros_tool_bandwidth_test

Needs a remote bandwidth-server. Skipped (long-running, network-dependent).

## Example Usage

```terraform
resource "routeros_tool_bandwidth_test" "bandwidth_test_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

