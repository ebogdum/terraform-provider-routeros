---
subcategory: "Bridge"
page_title: "RouterOS: routeros_interface_bridge"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_bridge

Manages the RouterOS `/interface/bridge` menu.

## Example Usage

```terraform
data "routeros_interface_bridge" "bridge_example" {
  # router   = "my-router"  # omit for the default router
  # filter   = { name = "some-name" }
  # proplist = ["name", "address"]
}
```

## Argument Reference

This data source supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to query.
* `filter` - (Optional) Map of field=value pairs to narrow the result set.
* `proplist` - (Optional) List of property names to project; smaller payload.
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

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `active_role` - Type: `enum(primary|secondary)`.
* `mac_address` - Type: `string`.
* `state` - Type: `string`.

