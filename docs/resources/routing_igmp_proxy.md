---
subcategory: "Routing"
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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `query_interval` - (Optional) Type: `string`.
* `query_response_interval` - (Optional) Type: `string`.
* `quick_leave` - (Optional) Type: `bool`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_routing_igmp_proxy.this 'home'
```
