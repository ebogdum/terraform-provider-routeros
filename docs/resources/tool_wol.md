---
page_title: "RouterOS: routeros_tool_wol"
description: |-
  RouterOS resource.
---

# Resource: routeros_tool_wol

Manages the RouterOS `/tool/wol` menu.

## Example Usage

```terraform
resource "routeros_tool_wol" "wol_example" {
  # router = "my-router"  # which router to target; omit for the default
  mac = "02:00:00:00:00:01"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `mac` - (Required) Type: `string`. Default: `02:00:00:00:00:01`.

