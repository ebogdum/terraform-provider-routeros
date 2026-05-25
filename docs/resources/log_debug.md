---
page_title: "RouterOS: routeros_log_debug"
description: |-
  RouterOS resource.
---

# Resource: routeros_log_debug

Manages the RouterOS `/log/debug` menu.

## Example Usage

```terraform
resource "routeros_log_debug" "debug_example" {
  # router = "my-router"  # which router to target; omit for the default
  message = "hello from terraform"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `message` - (Required) Type: `string`. Default: `tf-acc-test`.

