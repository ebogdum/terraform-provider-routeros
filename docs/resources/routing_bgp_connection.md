---
subcategory: "BGP"
page_title: "RouterOS: routeros_routing_bgp_connection"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Resource: routeros_routing_bgp_connection

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
resource "routeros_routing_bgp_connection" "connection_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # afi = "replace-me"
  # allow_as_in = "replace-me"
  # as = "replace-me"
  # as_override = "replace-me"
  # cisco_vpls_nlri_length_format = "replace-me"
  # connect = "replace-me"
  # default_originate = "replace-me"
  # default_prepend = "replace-me"
  # hold_time = "replace-me"
  # ignore_as_path_length = "replace-me"
  # input_accept_communities = "replace-me"
  # input_accept_ext_communities = "replace-me"
  # input_accept_large_communities = "replace-me"
  # input_accept_nlri = "replace-me"
  # input_affinity = "replace-me"
  # input_filter = "replace-me"
  # input_filter_communities = "replace-me"
  # input_filter_ext_communities = "replace-me"
  # input_filter_large_communities = "replace-me"
  # input_filter_unknown = "replace-me"
  # instance = "replace-me"
  # keep_sent_attributes = "replace-me"
  # keepalive_time = "replace-me"
  # listen = "replace-me"
  # local_address = "10.99.0.1"
  # local_port = "443"
  # local_role = "ibgp"
  # multihop = "replace-me"
  # name = "tf-example"
  # network_blackhole = "replace-me"
  # nexthop_choice = "replace-me"
  # no_client_to_client_reflection = "replace-me"
  # no_early_cut = "replace-me"
  # output_affinity = "replace-me"
  # output_filter = "replace-me"
  # output_network = "replace-me"
  # output_redistribute = "replace-me"
  # output_selection_policy = "replace-me"
  # remote_address = "10.99.0.1"
  # remote_allow_as = "replace-me"
  # remote_as = "replace-me"
  # remote_port = "443"
  # remove_private_as = "replace-me"
  # router_id = "replace-me"
  # routing_table = "main"
  # rx_min_ttl = "replace-me"
  # tcp_md5_key = "REDACTED"
  # template = "replace-me"
  # tx_ttl = "replace-me"
  # use_bfd = "replace-me"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `afi` - (Optional) Type: `string`.
* `allow_as_in` - (Read-only) Type: `string`.
* `as` - (Optional) Type: `string`.
* `as_override` - (Read-only) Type: `string`.
* `cisco_vpls_nlri_len_fmt` - (Optional) Type: `string`. RouterOS `cisco-vpls-nlri-len-fmt`.
* `cisco_vpls_nlri_length_format` - (Read-only) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connect` - (Optional) Type: `string`.
* `default_originate` - (Read-only) Type: `string`.
* `default_prepend` - (Read-only) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dynamic` - (Read-only) Type: `bool`.
* `hold_time` - (Optional) Type: `string`.
* `ignore_as_path_length` - (Read-only) Type: `string`.
* `input_accept_communities` - (Optional) Type: `string`.
* `input_accept_ext_communities` - (Optional) Type: `string`.
* `input_accept_large_communities` - (Optional) Type: `string`.
* `input_accept_nlri` - (Optional) Type: `string`.
* `input_add_path` - (Optional) Type: `string`. RouterOS `input.add-path`.
* `input_affinity` - (Optional) Type: `string`.
* `input_allow_as` - (Optional) Type: `string`. RouterOS `input.allow-as`.
* `input_attr_error_handling` - (Optional) Type: `string`. RouterOS `input.attr-error-handling`.
* `input_filter` - (Optional) Type: `string`.
* `input_filter_communities` - (Optional) Type: `string`.
* `input_filter_ext_communities` - (Optional) Type: `string`.
* `input_filter_large_communities` - (Optional) Type: `string`.
* `input_filter_nlri` - (Optional) Type: `string`. RouterOS `input.filter-nlri`.
* `input_filter_unknown` - (Optional) Type: `string`.
* `input_limit_process_routes_ipv4` - (Optional) Type: `string`. RouterOS `input.limit-process-routes-ipv4`.
* `input_limit_process_routes_ipv6` - (Optional) Type: `string`. RouterOS `input.limit-process-routes-ipv6`.
* `instance` - (Optional) Type: `string`.
* `invalid` - (Read-only) Type: `bool`.
* `keep_sent_attributes` - (Read-only) Type: `string`.
* `keepalive_time` - (Optional) Type: `string`.
* `listen` - (Optional) Type: `string`.
* `local_address` - (Optional) Type: `string`.
* `local_port` - (Optional) Type: `string`.
* `local_role` - (Optional) Type: `string`.
* `local_ttl` - (Optional) Type: `string`. RouterOS `local.ttl`.
* `multihop` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `network_blackhole` - (Read-only) Type: `string`.
* `nexthop_choice` - (Optional) Type: `string`.
* `no_client_to_client_reflection` - (Read-only) Type: `string`.
* `no_early_cut` - (Read-only) Type: `string`.
* `output_add_path` - (Optional) Type: `string`. RouterOS `output.add-path`.
* `output_affinity` - (Optional) Type: `string`.
* `output_as_override` - (Optional) Type: `string`. RouterOS `output.as-override`.
* `output_default_originate` - (Optional) Type: `string`. RouterOS `output.default-originate`.
* `output_default_prepend` - (Optional) Type: `string`. RouterOS `output.default-prepend`.
* `output_filter` - (Read-only) Type: `string`.
* `output_filter_chain` - (Optional) Type: `string`. RouterOS `output.filter-chain`.
* `output_filter_select` - (Optional) Type: `string`. RouterOS `output.filter-select`.
* `output_keep_sent_attributes` - (Optional) Type: `string`. RouterOS `output.keep-sent-attributes`.
* `output_network` - (Optional) Type: `string`.
* `output_network_blackhole` - (Optional) Type: `string`. RouterOS `output.network-blackhole`.
* `output_no_client_to_client_reflection` - (Optional) Type: `string`. RouterOS `output.no-client-to-client-reflection`.
* `output_no_early_cut` - (Optional) Type: `string`. RouterOS `output.no-early-cut`.
* `output_redistribute` - (Optional) Type: `string`.
* `output_remove_private_as` - (Optional) Type: `string`. RouterOS `output.remove-private-as`.
* `output_selection_policy` - (Read-only) Type: `string`.
* `remote_address` - (Optional) Type: `string`.
* `remote_allow_as` - (Read-only) Type: `string`.
* `remote_allowed_as` - (Optional) Type: `string`. RouterOS `remote.allowed-as`.
* `remote_as` - (Optional) Type: `string`.
* `remote_port` - (Optional) Type: `string`.
* `remote_ttl` - (Optional) Type: `string`. RouterOS `remote.ttl`.
* `remove_private_as` - (Read-only) Type: `string`.
* `router_id` - (Read-only) Type: `string`.
* `routing_table` - (Optional) Type: `string`.
* `rx_min_ttl` - (Read-only) Type: `string`.
* `save_to` - (Optional) Type: `string`. RouterOS `save-to`.
* `tcp_md5_key` - (Optional) Type: `string`. **Sensitive.**
* `template` - (Read-only) Type: `string`.
* `templates` - (Optional) Type: `string`. RouterOS `templates`.
* `tx_ttl` - (Read-only) Type: `string`.
* `use_bfd` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_bgp_connection.example '*3'

# Named router
terraform import routeros_routing_bgp_connection.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_bgp_connection.example 'home/my-resource-name'
```
