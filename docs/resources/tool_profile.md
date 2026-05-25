---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_profile"
description: |-
  Long-running CPU profiler.
---

# Resource: routeros_tool_profile

Long-running CPU profiler.

## Example Usage

```terraform
resource "routeros_tool_profile" "profile_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

