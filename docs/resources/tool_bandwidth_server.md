---
page_title: "RouterOS: routeros_tool_bandwidth_server"
description: |-
  RouterOS resource.
---

# Resource: routeros_tool_bandwidth_server

Manages the RouterOS `/tool/bandwidth-server` menu.

## Example Usage

```terraform
resource "routeros_tool_bandwidth_server" "bandwidth_server_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # allocate_udp_ports_from = 0
  # allowed_addresses4 = "replace-me"
  # allowed_addresses6 = "replace-me"
  # authenticate = false
  # enabled = false
  # max_sessions = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `allocate_udp_ports_from` - (Optional) Type: `int`. Beginning of UDP port range.
* `allowed_addresses4` - (Optional) Type: `string`.
* `allowed_addresses6` - (Optional) Type: `string`.
* `authenticate` - (Optional) Type: `bool`. Communicate only with authenticated clients.
* `enabled` - (Optional) Type: `bool`. Defines whether bandwidth server is enabled or not.
* `max_sessions` - (Optional) Type: `int`. Maximal simultaneous test count.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_tool_bandwidth_server.this 'home'
```
