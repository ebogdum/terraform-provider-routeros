---
subcategory: "IP"
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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `enabled` - (Optional) Type: `string`.
* `generic_timeout` - (Optional) Type: `duration`.
* `icmp_timeout` - (Optional) Type: `duration`.
* `liberal_tcp_tracking` - (Optional) Type: `bool`.
* `loose_tcp_tracking` - (Optional) Type: `bool`.
* `tcp_close_timeout` - (Optional) Type: `duration`.
* `tcp_close_wait_timeout` - (Optional) Type: `duration`.
* `tcp_established_timeout` - (Optional) Type: `duration`.
* `tcp_fin_wait_timeout` - (Optional) Type: `duration`.
* `tcp_last_ack_timeout` - (Optional) Type: `duration`.
* `tcp_max_retrans_timeout` - (Optional) Type: `duration`.
* `tcp_syn_received_timeout` - (Optional) Type: `duration`.
* `tcp_syn_sent_timeout` - (Optional) Type: `duration`.
* `tcp_time_wait_timeout` - (Optional) Type: `duration`.
* `tcp_unacked_timeout` - (Optional) Type: `duration`.
* `udp_stream_timeout` - (Optional) Type: `duration`.
* `udp_timeout` - (Optional) Type: `duration`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `active_ipv4` - Type: `bool`.
* `active_ipv6` - Type: `bool`.
* `max_entries` - Type: `int`.
* `total_entries` - Type: `int`.
* `total_ip4_entries` - Type: `int`.
* `total_ip6_entries` - Type: `int`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_firewall_connection_tracking.this 'home'
```
