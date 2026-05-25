---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_bridge"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_bridge

Manages the RouterOS `/interface/bridge` menu.

## Example Usage

```terraform
resource "routeros_interface_bridge" "bridge_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # ageing_time = "1h"
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # auto_mac = false
  # dhcp_snooping = false
  # fast_forward = false
  # forward_delay = "1h"
  # igmp_snooping = false
  # max_learned_entries = "replace-me"
  # max_message_age = "1h"
  # mlag_heartbeat = "1h"
  # mlag_peer_port = "443"
  # mlag_priority = 0
  # mtu = "replace-me"
  # name = "example"
  # port_cost_mode = "replace-me"
  # priority = 0
  # protocol_mode = "replace-me"
  # ra_guard = false
  # transmit_hold_count = 0
  # vlan_filtering = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `ageing_time` - (Optional) Type: `duration`.
* `arp` - (Optional) Type: `string`.
* `arp_timeout` - (Optional) Type: `string`.
* `auto_mac` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dhcp_snooping` - (Optional) Type: `bool`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `fast_forward` - (Optional) Type: `bool`.
* `forward_delay` - (Optional) Type: `duration`.
* `igmp_snooping` - (Optional) Type: `bool`.
* `max_learned_entries` - (Optional) Type: `string`.
* `max_message_age` - (Optional) Type: `duration`.
* `mlag_heartbeat` - (Optional) Type: `duration`.
* `mlag_peer_port` - (Optional) Type: `string`.
* `mlag_priority` - (Optional) Type: `int`.
* `mtu` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `port_cost_mode` - (Optional) Type: `string`.
* `priority` - (Optional) Type: `int`.
* `protocol_mode` - (Optional) Type: `string`.
* `ra_guard` - (Optional) Type: `bool`.
* `transmit_hold_count` - (Optional) Type: `int`.
* `vlan_filtering` - (Optional) Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_bridge.example '*3'

# Named router
terraform import routeros_interface_bridge.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_bridge.example 'home/my-resource-name'
```
