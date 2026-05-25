---
subcategory: "DNS"
page_title: "RouterOS: routeros_ip_dns"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_dns

Manages the RouterOS `/ip/dns` menu.

## Example Usage

```terraform
resource "routeros_ip_dns" "dns_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # address_list_extra_time = "1h"
  # allow_remote_requests = false
  # cache_max_ttl = "1h"
  # cache_size = 0
  # doh_max_concurrent_queries = 0
  # doh_max_server_connections = 0
  # doh_timeout = "1h"
  # max_concurrent_queries = 0
  # max_concurrent_tcp_sessions = 0
  # max_udp_packet_size = 0
  # mdns_repeat_ifaces = "replace-me"
  # query_server_timeout = "1h"
  # query_total_timeout = "1h"
  # servers = "replace-me"
  # use_doh_server = "replace-me"
  # verify_doh_cert = false
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address_list_extra_time` - (Optional) Type: `duration`.
* `allow_remote_requests` - (Optional) Type: `bool`.
* `cache_max_ttl` - (Optional) Type: `duration`.
* `cache_size` - (Optional) Type: `int`.
* `doh_max_concurrent_queries` - (Optional) Type: `int`.
* `doh_max_server_connections` - (Optional) Type: `int`.
* `doh_timeout` - (Optional) Type: `duration`.
* `max_concurrent_queries` - (Optional) Type: `int`.
* `max_concurrent_tcp_sessions` - (Optional) Type: `int`.
* `max_udp_packet_size` - (Optional) Type: `int`.
* `mdns_repeat_ifaces` - (Optional) Type: `string`.
* `query_server_timeout` - (Optional) Type: `duration`.
* `query_total_timeout` - (Optional) Type: `duration`.
* `servers` - (Optional) Type: `string`.
* `use_doh_server` - (Optional) Type: `string`.
* `verify_doh_cert` - (Optional) Type: `bool`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `cache_used` - Type: `int`.
* `dynamic_servers` - Type: `ip`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_dns.this 'home'
```
