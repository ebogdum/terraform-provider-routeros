---
page_title: "RouterOS: routeros_ipv6_settings"
description: |-
  RouterOS resource.
---

# Resource: routeros_ipv6_settings

Manages the RouterOS `/ipv6/settings` menu.

## Example Usage

```terraform
resource "routeros_ipv6_settings" "settings_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # accept_redirects = "replace-me"
  # accept_router_advertisements = "replace-me"
  # accept_router_advertisements_on = "replace-me"
  # allow_fast_path = false
  # disable_ipv6 = false
  # disable_link_local_address = "10.99.0.0/24"
  # forward = false
  # max_neighbor_entries = 0
  # min_neighbor_entries = 0
  # multipath_hash_policy = "replace-me"
  # soft_max_neighbor_entries = 0
  # stale_neighbor_detect_interval = 0
  # stale_neighbor_timeout = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `accept_redirects` - (Optional) Type: `string`.
* `accept_router_advertisements` - (Optional) Type: `string`.
* `accept_router_advertisements_on` - (Optional) Type: `string`.
* `allow_fast_path` - (Optional) Type: `bool`.
* `disable_ipv6` - (Optional) Type: `bool`.
* `disable_link_local_address` - (Optional) Type: `bool`.
* `forward` - (Optional) Type: `bool`.
* `max_neighbor_entries` - (Optional) Type: `int`.
* `min_neighbor_entries` - (Optional) Type: `int`.
* `multipath_hash_policy` - (Optional) Type: `string`.
* `soft_max_neighbor_entries` - (Optional) Type: `int`.
* `stale_neighbor_detect_interval` - (Optional) Type: `int`.
* `stale_neighbor_timeout` - (Optional) Type: `int`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `ipv6_fast_path_active` - Type: `bool`.
* `ipv6_fast_path_bytes` - Type: `int`.
* `ipv6_fast_path_packets` - Type: `int`.
* `ipv6_fasttrack_active` - Type: `bool`.
* `ipv6_fasttrack_bytes` - Type: `int`.
* `ipv6_fasttrack_packets` - Type: `int`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_ipv6_settings.this 'home'
```
