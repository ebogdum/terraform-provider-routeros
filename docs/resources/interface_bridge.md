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
  # admin_mac = "replace-me"
  # ageing_time = "1h"
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # auto_mac = false
  # dhcp_snooping = false
  # ether_type = "replace-me"
  # fast_forward = false
  # forward_delay = "1h"
  # frame_types = "replace-me"
  # igmp_snooping = false
  # ingress_filtering = "replace-me"
  # max_learned_entries = "replace-me"
  # max_message_age = "1h"
  # mlag_heartbeat = "1h"
  # mlag_peer_port = "443"
  # mlag_priority = 0
  # mtu = "replace-me"
  # mvrp = "replace-me"
  # name = "tf-example"
  # port_cost_mode = "replace-me"
  # priority = 0
  # protocol_mode = "replace-me"
  # pvid = "replace-me"
  # ra_guard = false
  # region_name = "replace-me"
  # region_revision = "replace-me"
  # transmit_hold_count = 0
  # vlan_filtering = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `admin_mac` - (Optional) Type: `string`.
* `ageing_time` - (Optional) Type: `duration`.
* `arp` - (Optional) Type: `string`.
* `arp_timeout` - (Optional) Type: `string`.
* `auto_mac` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dhcp_snooping` - (Optional) Type: `bool`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `ether_type` - (Optional) Type: `string`.
* `fast_forward` - (Optional) Type: `bool`.
* `forward_delay` - (Optional) Type: `duration`.
* `frame_types` - (Optional) Type: `string`.
* `igmp_snooping` - (Optional) Type: `bool`.
* `ingress_filtering` - (Optional) Type: `string`.
* `max_learned_entries` - (Optional) Type: `string`.
* `max_message_age` - (Optional) Type: `duration`.
* `mlag_heartbeat` - (Optional) Type: `duration`.
* `mlag_peer_port` - (Optional) Type: `string`.
* `mlag_priority` - (Optional) Type: `int`.
* `mtu` - (Optional) Type: `string`.
* `mvrp` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `port_cost_mode` - (Optional) Type: `string`.
* `priority` - (Optional) Type: `int`.
* `protocol_mode` - (Optional) Type: `string`.
* `pvid` - (Optional) Type: `string`.
* `ra_guard` - (Optional) Type: `bool`.
* `region_name` - (Optional) Type: `string`.
* `region_revision` - (Optional) Type: `string`.
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
