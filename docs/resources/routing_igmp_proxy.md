---
page_title: "RouterOS: routeros_routing_igmp_proxy"
description: |-
  RouterOS resource.
---

# Resource: routeros_routing_igmp_proxy

Manages the RouterOS `/routing/igmp-proxy` menu.

## Example Usage

```terraform
resource "routeros_routing_igmp_proxy" "igmp_proxy_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # query_interval = "1h"
  # query_response_interval = "1h"
  # quick_leave = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `query_interval` - (Optional) Type: `duration`.
* `query_response_interval` - (Optional) Type: `duration`.
* `quick_leave` - (Optional) Type: `bool`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_routing_igmp_proxy.this 'home'
```
