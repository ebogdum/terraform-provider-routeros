---
subcategory: "BGP"
page_title: "RouterOS: routeros_routing_bgp_connection"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Data Source: routeros_routing_bgp_connection

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
data "routeros_routing_bgp_connection" "connection_example" {
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
* `afi` - (Optional) Type: `string`.
* `allow_as_in` - (Optional) Type: `string`.
* `as` - (Optional) Type: `string`.
* `as_override` - (Optional) Type: `string`.
* `cisco_vpls_nlri_length_format` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connect` - (Optional) Type: `string`.
* `default_originate` - (Optional) Type: `string`.
* `default_prepend` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `hold_time` - (Optional) Type: `string`.
* `ignore_as_path_length` - (Optional) Type: `string`.
* `input_accept_communities` - (Optional) Type: `string`.
* `input_accept_ext_communities` - (Optional) Type: `string`.
* `input_accept_large_communities` - (Optional) Type: `string`.
* `input_accept_nlri` - (Optional) Type: `string`.
* `input_affinity` - (Optional) Type: `string`.
* `input_filter` - (Optional) Type: `string`.
* `input_filter_communities` - (Optional) Type: `string`.
* `input_filter_ext_communities` - (Optional) Type: `string`.
* `input_filter_large_communities` - (Optional) Type: `string`.
* `input_filter_unknown` - (Optional) Type: `string`.
* `instance` - (Optional) Type: `string`.
* `keep_sent_attributes` - (Optional) Type: `string`.
* `keepalive_time` - (Optional) Type: `string`.
* `listen` - (Optional) Type: `string`.
* `local_address` - (Optional) Type: `string`.
* `local_port` - (Optional) Type: `string`.
* `local_role` - (Optional) Type: `enum(ibgp|ibgp-rr|ebgp|ebgp-provider|ebgp-rs|ebgp-rs-client, ...)`.
* `multihop` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `network_blackhole` - (Optional) Type: `string`.
* `nexthop_choice` - (Optional) Type: `string`.
* `no_client_to_client_reflection` - (Optional) Type: `string`.
* `no_early_cut` - (Optional) Type: `string`.
* `output_affinity` - (Optional) Type: `string`.
* `output_filter` - (Optional) Type: `string`.
* `output_network` - (Optional) Type: `string`.
* `output_redistribute` - (Optional) Type: `string`.
* `output_selection_policy` - (Optional) Type: `string`.
* `remote_address` - (Optional) Type: `string`.
* `remote_allow_as` - (Optional) Type: `string`.
* `remote_as` - (Optional) Type: `string`.
* `remote_port` - (Optional) Type: `string`.
* `remove_private_as` - (Optional) Type: `string`.
* `router_id` - (Optional) Type: `string`.
* `routing_table` - (Optional) Type: `string`.
* `rx_min_ttl` - (Optional) Type: `string`.
* `tcp_md5_key` - (Optional) Type: `string`. **Sensitive.**
* `template` - (Optional) Type: `string`.
* `tx_ttl` - (Optional) Type: `string`.
* `use_bfd` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.

