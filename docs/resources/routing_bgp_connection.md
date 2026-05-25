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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
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

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.

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
