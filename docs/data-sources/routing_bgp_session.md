---
page_title: "RouterOS: routeros_routing_bgp_session"
description: |-
  RouterOS resource.
---

# Data Source: routeros_routing_bgp_session

Manages the RouterOS `/routing/bgp/session` menu.

## Example Usage

```terraform
data "routeros_routing_bgp_session" "session_example" {
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
* `allow_as_in` - (Optional) Type: `int`.
* `as_override` - (Optional) Type: `bool`.
* `cisco_vpls_nlri_length_format` - (Optional) Type: `enum(auto bits|auto bytes|bits|bytes)`.
* `clear` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default_originate` - (Optional) Type: `enum(never|if installed|always)`.
* `default_prepend` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dump_adv` - (Optional) Type: `string`.
* `ebgp` - (Optional) Type: `bool`.
* `established` - (Optional) Type: `bool`.
* `hold_time` - (Optional) Type: `duration`.
* `ibgp` - (Optional) Type: `bool`.
* `input_affinity` - (Optional) Type: `enum(remote as|vrf|afi|instance|alone|main)`.
* `input_filter` - (Optional) Type: `string`.
* `instance` - (Optional) Type: `string`.
* `keep_sent_attributes` - (Optional) Type: `bool`.
* `keepalive_time` - (Optional) Type: `duration`.
* `limit_exceeded` - (Optional) Type: `bool`.
* `local_address` - (Optional) Type: `string`.
* `local_afi` - (Optional) Type: `string`. Default: `2`.
* `local_capabilities` - (Optional) Type: `string`.
* `local_cluster_id` - (Optional) Type: `ip`.
* `local_eor` - (Optional) Type: `string`.
* `local_id` - (Optional) Type: `ip`.
* `local_last_notification` - (Optional) Type: `string`.
* `local_role` - (Optional) Type: `enum(ibgp|ibgp rr|ebgp|ebgp provider|ebgp rs|ebgp rs client, ...)`.
* `multihop` - (Optional) Type: `bool`.
* `name` - (Optional) Type: `string`.
* `nexthop_choice` - (Optional) Type: `enum(default|force self|propagate)`.
* `no_client_to_client_reflection` - (Optional) Type: `bool`.
* `no_early_cut` - (Optional) Type: `bool`.
* `output_affinity` - (Optional) Type: `string`.
* `output_filter` - (Optional) Type: `string`.
* `output_network` - (Optional) Type: `string`.
* `output_selection_policy` - (Optional) Type: `string`.
* `prefix_count` - (Optional) Type: `int`.
* `refresh` - (Optional) Type: `string`.
* `remote_address` - (Optional) Type: `string`.
* `remote_afi` - (Optional) Type: `string`. Default: `2`.
* `remote_as` - (Optional) Type: `string`.
* `remote_capabilities` - (Optional) Type: `string`.
* `remote_eor` - (Optional) Type: `string`.
* `remote_gr_afi` - (Optional) Type: `string`.
* `remote_gr_afi_fw_path` - (Optional) Type: `string`.
* `remote_gr_restart` - (Optional) Type: `bool`.
* `remote_gr_time` - (Optional) Type: `int`.
* `remote_hold_time` - (Optional) Type: `duration`.
* `remote_id` - (Optional) Type: `ip`.
* `remote_last_notification` - (Optional) Type: `string`.
* `remote_refused_capability` - (Optional) Type: `bool`.
* `remove_private_as` - (Optional) Type: `bool`.
* `resend` - (Optional) Type: `string`.
* `routing_table` - (Optional) Type: `string`.
* `stop` - (Optional) Type: `string`.
* `stopped` - (Optional) Type: `bool`.
* `tx_rx_bytes` - (Optional) Type: `string`.
* `tx_rx_messages` - (Optional) Type: `string`.
* `use_bfd` - (Optional) Type: `bool`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `uptime` - Type: `string`.

