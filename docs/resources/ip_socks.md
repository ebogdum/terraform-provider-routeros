---
page_title: "RouterOS: routeros_ip_socks"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_socks

Manages the RouterOS `/ip/socks` menu.

## Example Usage

```terraform
resource "routeros_ip_socks" "socks_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # auth_method = "replace-me"
  # connection_idle_timeout = "1h"
  # enabled = false
  # max_connections = 0
  # port = "443"
  # version = 0
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `auth_method` - (Optional) Type: `string`.
* `connection_idle_timeout` - (Optional) Type: `duration`.
* `enabled` - (Optional) Type: `bool`.
* `max_connections` - (Optional) Type: `int`.
* `port` - (Optional) Type: `int`.
* `version` - (Optional) Type: `int`.
* `vrf` - (Optional) Type: `string`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_socks.this 'home'
```
