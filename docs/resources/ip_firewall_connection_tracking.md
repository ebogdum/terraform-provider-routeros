---
subcategory: "Firewall"
page_title: "RouterOS: routeros_ip_firewall_connection_tracking"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_firewall_connection_tracking

Manages the RouterOS `/ip/firewall/connection/tracking` menu.

## Example Usage

```terraform
resource "routeros_ip_firewall_connection_tracking" "tracking_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # enabled = "replace-me"
  # generic_timeout = "1h"
  # icmp_timeout = "1h"
  # liberal_tcp_tracking = false
  # loose_tcp_tracking = false
  # tcp_close_timeout = "1h"
  # tcp_close_wait_timeout = "1h"
  # tcp_established_timeout = "1h"
  # tcp_fin_wait_timeout = "1h"
  # tcp_last_ack_timeout = "1h"
  # tcp_max_retrans_timeout = "1h"
  # tcp_syn_received_timeout = "1h"
  # tcp_syn_sent_timeout = "1h"
  # tcp_time_wait_timeout = "1h"
  # tcp_unacked_timeout = "1h"
  # udp_stream_timeout = "1h"
  # udp_timeout = "1h"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `active_ipv4` - (Optional) Type: `bool`.
* `active_ipv6` - (Optional) Type: `bool`.
* `enabled` - (Optional) Type: `string`.
* `generic_timeout` - (Optional) Type: `string`.
* `icmp_timeout` - (Optional) Type: `string`.
* `liberal_tcp_tracking` - (Optional) Type: `bool`.
* `loose_tcp_tracking` - (Optional) Type: `bool`.
* `max_entries` - (Optional) Type: `int`.
* `tcp_close_timeout` - (Optional) Type: `string`.
* `tcp_close_wait_timeout` - (Optional) Type: `string`.
* `tcp_established_timeout` - (Optional) Type: `string`.
* `tcp_fin_wait_timeout` - (Optional) Type: `string`.
* `tcp_last_ack_timeout` - (Optional) Type: `string`.
* `tcp_max_retrans_timeout` - (Optional) Type: `string`.
* `tcp_syn_received_timeout` - (Optional) Type: `string`.
* `tcp_syn_sent_timeout` - (Optional) Type: `string`.
* `tcp_time_wait_timeout` - (Optional) Type: `string`.
* `tcp_unacked_timeout` - (Optional) Type: `string`.
* `total_entries` - (Optional) Type: `int`.
* `total_ip4_entries` - (Optional) Type: `int`.
* `total_ip6_entries` - (Optional) Type: `int`.
* `udp_stream_timeout` - (Optional) Type: `string`.
* `udp_timeout` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_firewall_connection_tracking.this 'home'
```
