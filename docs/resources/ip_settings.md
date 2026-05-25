---
page_title: "RouterOS: routeros_ip_settings"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_settings

Manages the RouterOS `/ip/settings` menu.

## Example Usage

```terraform
resource "routeros_ip_settings" "settings_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # accept_redirects = false
  # accept_source_route = false
  # allow_fast_path = false
  # arp_timeout = "1h"
  # icmp_errors_use_inbound_interface_address = "10.99.0.0/24"
  # icmp_rate_limit = 0
  # icmp_rate_mask = 0
  # ip_forward = false
  # ipv4_multipath_hash_policy = "replace-me"
  # max_neighbor_entries = 0
  # rp_filter = false
  # secure_redirects = false
  # send_redirects = false
  # tcp_syncookies = false
  # tcp_timestamps = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `accept_redirects` - (Optional) Type: `bool`.
* `accept_source_route` - (Optional) Type: `bool`.
* `allow_fast_path` - (Optional) Type: `bool`.
* `arp_timeout` - (Optional) Type: `duration`.
* `icmp_errors_use_inbound_interface_address` - (Optional) Type: `bool`.
* `icmp_rate_limit` - (Optional) Type: `int`.
* `icmp_rate_mask` - (Optional) Type: `int`.
* `ip_forward` - (Optional) Type: `bool`.
* `ipv4_multipath_hash_policy` - (Optional) Type: `string`.
* `max_neighbor_entries` - (Optional) Type: `int`.
* `rp_filter` - (Optional) Type: `bool`.
* `secure_redirects` - (Optional) Type: `bool`.
* `send_redirects` - (Optional) Type: `bool`.
* `tcp_syncookies` - (Optional) Type: `bool`.
* `tcp_timestamps` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `ipv4_fast_path_active` - Type: `bool`.
* `ipv4_fast_path_bytes` - Type: `int`.
* `ipv4_fast_path_packets` - Type: `int`.
* `ipv4_fasttrack_active` - Type: `bool`.
* `ipv4_fasttrack_bytes` - Type: `int`.
* `ipv4_fasttrack_packets` - Type: `int`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_settings.this 'home'
```
