---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_proxy"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_proxy

Manages the RouterOS `/ip/proxy` menu.

## Example Usage

```terraform
resource "routeros_ip_proxy" "proxy_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # always_from_cache = false
  # anonymous = false
  # cache_administrator = "replace-me"
  # cache_hit_dscp = 0
  # cache_on_disk = false
  # cache_path = "replace-me"
  # enabled = false
  # max_cache_object_size = 0
  # max_cache_size = "replace-me"
  # max_client_connections = 0
  # max_fresh_time = "1h"
  # max_server_connections = 0
  # parent_proxy = "10.99.0.1"
  # parent_proxy_port = "443"
  # port = "443"
  # serialize_connections = false
  # src_address = "10.99.0.0/24"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `always_from_cache` - (Optional) Type: `bool`.
* `anonymous` - (Optional) Type: `bool`.
* `cache_administrator` - (Optional) Type: `string`.
* `cache_hit_dscp` - (Optional) Type: `int`.
* `cache_on_disk` - (Optional) Type: `bool`.
* `cache_path` - (Optional) Type: `string`.
* `enabled` - (Optional) Type: `bool`.
* `max_cache_object_size` - (Optional) Type: `int`.
* `max_cache_size` - (Optional) Type: `string`.
* `max_client_connections` - (Optional) Type: `int`.
* `max_fresh_time` - (Optional) Type: `string`.
* `max_server_connections` - (Optional) Type: `int`.
* `parent_proxy` - (Optional) Type: `string`.
* `parent_proxy_port` - (Optional) Type: `int`.
* `port` - (Optional) Type: `int`.
* `serialize_connections` - (Optional) Type: `bool`.
* `src_address` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_proxy.this 'home'
```
