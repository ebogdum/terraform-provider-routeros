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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `active_role` - (Read-only) Type: `string`.
* `add_dhcp_option_82` - (Read-only) Type: `bool`.
* `admin_mac` - (Optional) Type: `string`.
* `admin_mac_address` - (Read-only) Type: `string`.
* `ageing_time` - (Optional) Type: `string`.
* `arp` - (Optional) Type: `string`.
* `arp_timeout` - (Optional) Type: `string`.
* `auto_mac` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dhcp_agent_circuit_id` - (Optional) Type: `string`. RouterOS `dhcp-agent-circuit-id`.
* `dhcp_agent_remote_id` - (Optional) Type: `string`. RouterOS `dhcp-agent-remote-id`.
* `dhcp_snooping` - (Optional) Type: `bool`.
* `dhcpv6_agent_circuit_id` - (Optional) Type: `string`. RouterOS `dhcpv6-agent-circuit-id`.
* `dhcpv6_agent_remote_id` - (Optional) Type: `string`. RouterOS `dhcpv6-agent-remote-id`.
* `dhcpv6_snooping` - (Optional) Type: `string`. RouterOS `dhcpv6-snooping`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dumb` - (Read-only) Type: `string`.
* `ether_type` - (Optional) Type: `string`.
* `fast_forward` - (Optional) Type: `bool`.
* `forward_delay` - (Optional) Type: `string`.
* `forward_reserved` - (Read-only) Type: `bool`.
* `forward_reserved_addresses` - (Optional) Type: `string`. RouterOS `forward-reserved-addresses`.
* `fp_tx_rx_packet_rate` - (Read-only) Type: `string`.
* `fp_tx_rx_rate` - (Read-only) Type: `string`.
* `frame_types` - (Optional) Type: `string`.
* `heartbeat` - (Read-only) Type: `string`.
* `igmp` - (Read-only) Type: `string`.
* `igmp_snooping` - (Optional) Type: `bool`.
* `igmp_version` - (Optional) Type: `string`.
* `ingress_filtering` - (Optional) Type: `bool`.
* `last_member_interval` - (Optional) Type: `string`.
* `last_member_query_count` - (Optional) Type: `int`.
* `mac_address` - (Read-only) Type: `string`.
* `max_hops` - (Optional) Type: `int`.
* `max_learned_entries` - (Optional) Type: `string`.
* `max_message_age` - (Optional) Type: `string`.
* `membership_interval` - (Optional) Type: `string`.
* `mlag_heartbeat` - (Optional) Type: `string`.
* `mlag_peer_port` - (Optional) Type: `string`.
* `mlag_priority` - (Optional) Type: `int`.
* `mld_version` - (Optional) Type: `string`.
* `mstp` - (Read-only) Type: `string`.
* `mtu` - (Optional) Type: `string`. Bridge MTU. A number, or `auto` (the default) to follow the smallest port MTU.
* `multicast_querier` - (Optional) Type: `bool`.
* `multicast_router` - (Optional) Type: `string`.
* `mvrp` - (Optional) Type: `bool`.
* `name` - (Optional) Type: `string`.
* `peer_port` - (Read-only) Type: `string`.
* `port_cost_mode` - (Optional) Type: `string`.
* `priority` - (Optional) Type: `int`.
* `protocol_mode` - (Optional) Type: `string`.
* `pvid` - (Optional) Type: `int`.
* `querier_interval` - (Optional) Type: `string`.
* `query_interval` - (Optional) Type: `string`.
* `query_response_interval` - (Optional) Type: `string`.
* `ra_guard` - (Optional) Type: `bool`.
* `region_name` - (Optional) Type: `string`.
* `region_revision` - (Optional) Type: `int`.
* `startup_query_count` - (Optional) Type: `int`.
* `startup_query_interval` - (Optional) Type: `string`.
* `state` - (Read-only) Type: `string`.
* `status` - (Read-only) Type: `int`.
* `transmit_hold_count` - (Optional) Type: `int`.
* `tx_rx_packet_rate` - (Read-only) Type: `string`.
* `tx_rx_rate` - (Read-only) Type: `string`.
* `type` - (Read-only) Type: `string`.
* `vlan` - (Read-only) Type: `string`.
* `vlan_filtering` - (Optional) Type: `bool`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
