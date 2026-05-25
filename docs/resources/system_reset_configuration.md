---
page_title: "RouterOS: routeros_system_reset_configuration"
description: |-
  Resets RouterOS to defaults. Curl confirms RouterOS accepts the POST and
---

# Resource: routeros_system_reset_configuration

Resets RouterOS to defaults. Curl confirms RouterOS accepts the POST and
returns 200, but immediately afterwards the router drops the IP config and
cannot finish the test framework's destroy step (connection times out).
Skipped from automated acc tests; verified manually.


## Example Usage

```terraform
resource "routeros_system_reset_configuration" "reset_configuration_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

