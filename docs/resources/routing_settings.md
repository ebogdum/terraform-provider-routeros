---
subcategory: "Routing"
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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `check_gateway_ping_count` - (Optional) Type: `int`.
* `check_gateway_ping_interval` - (Optional) Type: `string`.
* `check_gateway_ping_timeout` - (Optional) Type: `string`.
* `connected_in_chain` - (Optional) Type: `string`. RouterOS `connected-in-chain`.
* `dynamic_in_chain` - (Optional) Type: `string`. RouterOS `dynamic-in-chain`.
* `policy_rules` - (Optional) Type: `list`.
* `single_process` - (Optional) Type: `bool`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_routing_settings.this 'home'
```
