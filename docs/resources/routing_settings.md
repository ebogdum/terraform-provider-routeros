---
page_title: "RouterOS: routeros_routing_settings"
description: |-
  RouterOS resource.
---

# Resource: routeros_routing_settings

Manages the RouterOS `/routing/settings` menu.

## Example Usage

```terraform
resource "routeros_routing_settings" "settings_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # check_gateway_ping_count = 0
  # check_gateway_ping_interval = "1h"
  # check_gateway_ping_timeout = "1h"
  # policy_rules = []
  # single_process = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `check_gateway_ping_count` - (Optional) Type: `int`.
* `check_gateway_ping_interval` - (Optional) Type: `duration`.
* `check_gateway_ping_timeout` - (Optional) Type: `duration`.
* `policy_rules` - (Optional) Type: `list`.
* `single_process` - (Optional) Type: `bool`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_routing_settings.this 'home'
```
