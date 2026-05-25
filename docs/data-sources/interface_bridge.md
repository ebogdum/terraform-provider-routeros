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

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

