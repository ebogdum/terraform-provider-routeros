---
subcategory: "Bridge"
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
  # add_dhcp_option_82 = false
  # admin_mac = "replace-me"
  # admin_mac_address = "10.99.0.0/24"
  # ageing_time = "300"
  # arp = "enabled"
  # arp_timeout = "1h"
  # auto_mac = false
  # dhcp_snooping = false
  # dumb = "replace-me"
  # ether_type = "0x8100"
  # fast_forward = true
  # forward_delay = "15"
  # forward_reserved = false
  # fp_tx_rx_packet_rate = "replace-me"
  # fp_tx_rx_rate = "replace-me"
  # frame_types = "admit-all"
  # heartbeat = "5"
  # igmp = "replace-me"
  # igmp_snooping = false
  # igmp_version = "2"
  # ingress_filtering = true
  # last_member_interval = "100"
  # last_member_query_count = 2
  # max_hops = 20
  # max_learned_entries = "unlimited"
  # max_message_age = "20"
  # membership_interval = "26000"
  # mlag_heartbeat = "1h"
  # mlag_peer_port = "443"
  # mlag_priority = 0
  # mld_version = "1"
  # mstp = "replace-me"
  # mtu = 0
  # multicast_querier = false
  # multicast_router = "temporary-query"
  # mvrp = false
  # name = "tf-example"
  # peer_port = "443"
  # port_cost_mode = "long"
  # priority = 32768
  # protocol_mode = "rstp"
  # pvid = 0
  # querier_interval = "25500"
  # query_interval = "12500"
  # query_response_interval = "1000"
  # ra_guard = false
  # region_name = "replace-me"
  # region_revision = 0
  # startup_query_count = 2
  # startup_query_interval = "3125"
  # status = 0
  # transmit_hold_count = 6
  # tx_rx_packet_rate = "replace-me"
  # tx_rx_rate = "replace-me"
  # type = "replace-me"
  # vlan = "replace-me"
  # vlan_filtering = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `add_dhcp_option_82` - (Optional) Type: `bool`.
* `admin_mac` - (Optional) Type: `string`.
* `admin_mac_address` - (Optional) Type: `string`.
* `ageing_time` - (Optional) Type: `duration`. Default: `300`.
* `arp` - (Optional) Type: `enum(disabled|enabled|proxy-arp|reply-only|local-proxy-arp)`. Default: `1`.
* `arp_timeout` - (Optional) Type: `duration`.
* `auto_mac` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dhcp_snooping` - (Optional) Type: `bool`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dumb` - (Optional) Type: `string`.
* `ether_type` - (Optional) Type: `enum(0x8100|0x88a8|0x9100)`. Default: `33024`.
* `fast_forward` - (Optional) Type: `bool`. Default: `1`.
* `forward_delay` - (Optional) Type: `duration`. Default: `15`.
* `forward_reserved` - (Optional) Type: `bool`.
* `fp_tx_rx_packet_rate` - (Optional) Type: `string`.
* `fp_tx_rx_rate` - (Optional) Type: `string`.
* `frame_types` - (Optional) Type: `enum(admit-all|admit-only-vlan-tagged|admit-only-untagged-and-priority-tagged)`.
* `heartbeat` - (Optional) Type: `duration`. Default: `5`.
* `igmp` - (Optional) Type: `string`.
* `igmp_snooping` - (Optional) Type: `bool`.
* `igmp_version` - (Optional) Type: `enum(2|3)`. Default: `2`.
* `ingress_filtering` - (Optional) Type: `bool`. Default: `1`.
* `last_member_interval` - (Optional) Type: `string`. Default: `100`.
* `last_member_query_count` - (Optional) Type: `int`. Default: `2`.
* `max_hops` - (Optional) Type: `int`. Default: `20`.
* `max_learned_entries` - (Optional) Type: `enum(unlimited|auto)`. Default: `4.294967295e+09`.
* `max_message_age` - (Optional) Type: `duration`. Default: `20`.
* `membership_interval` - (Optional) Type: `string`. Default: `26000`.
* `mlag_heartbeat` - (Optional) Type: `duration`.
* `mlag_peer_port` - (Optional) Type: `string`.
* `mlag_priority` - (Optional) Type: `int`.
* `mld_version` - (Optional) Type: `enum(|1|2)`. Default: `1`.
* `mstp` - (Optional) Type: `string`.
* `mtu` - (Optional) Type: `int`. Default: `0`.
* `multicast_querier` - (Optional) Type: `bool`.
* `multicast_router` - (Optional) Type: `enum(disabled|temporary-query|permanent)`. Default: `1`.
* `mvrp` - (Optional) Type: `bool`.
* `name` - (Optional) Type: `string`.
* `peer_port` - (Optional) Type: `string`.
* `port_cost_mode` - (Optional) Type: `enum(short|long)`. Default: `1`.
* `priority` - (Optional) Type: `int`. Default: `32768`.
* `protocol_mode` - (Optional) Type: `enum(none|stp|rstp|mstp)`. Default: `2`.
* `pvid` - (Optional) Type: `int`.
* `querier_interval` - (Optional) Type: `string`. Default: `25500`.
* `query_interval` - (Optional) Type: `string`. Default: `12500`.
* `query_response_interval` - (Optional) Type: `string`. Default: `1000`.
* `ra_guard` - (Optional) Type: `bool`.
* `region_name` - (Optional) Type: `string`.
* `region_revision` - (Optional) Type: `int`.
* `startup_query_count` - (Optional) Type: `int`. Default: `2`.
* `startup_query_interval` - (Optional) Type: `string`. Default: `3125`.
* `status` - (Optional) Type: `int`.
* `transmit_hold_count` - (Optional) Type: `int`. Default: `6`.
* `tx_rx_packet_rate` - (Optional) Type: `string`.
* `tx_rx_rate` - (Optional) Type: `string`.
* `type` - (Optional) Type: `string`.
* `vlan` - (Optional) Type: `string`.
* `vlan_filtering` - (Optional) Type: `bool`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `active_role` - Type: `enum(primary|secondary)`.
* `mac_address` - Type: `string`.
* `state` - Type: `string`.

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
