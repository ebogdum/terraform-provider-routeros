---
subcategory: "BGP"
page_title: "RouterOS: routeros_routing_bgp_template"
description: |-
  RouterOS resource.
---

# Resource: routeros_routing_bgp_template

Manages the RouterOS `/routing/bgp/template` menu.

## Example Usage

```terraform
resource "routeros_routing_bgp_template" "template_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # afi = "replace-me"
  # allow_as_in = "replace-me"
  # as = "65000"
  # as_override = "replace-me"
  # cisco_vpls_nlri_length_format = "replace-me"
  # cluster_id = "replace-me"
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
  # keep_sent_attributes = "replace-me"
  # keepalive_time = "replace-me"
  # multihop = "replace-me"
  # nexthop_choice = "replace-me"
  # no_client_to_client_reflection = "replace-me"
  # no_early_cut = "replace-me"
  # output_affinity = "replace-me"
  # output_filter = "replace-me"
  # output_network = "replace-me"
  # output_redistribute = "replace-me"
  # output_selection_policy = "replace-me"
  # remove_private_as = "replace-me"
  # router_id = "1.1.1.1"
  # routing_table = "main"
  # templates = "replace-me"
  # use_bfd = "replace-me"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `afi` - (Optional) Type: `string`.
* `allow_as_in` - (Optional) Type: `string`.
* `as` - (Optional) Type: `string`. Default: `65000`.
* `as_override` - (Optional) Type: `string`.
* `cisco_vpls_nlri_length_format` - (Optional) Type: `string`.
* `cluster_id` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
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
* `keep_sent_attributes` - (Optional) Type: `string`.
* `keepalive_time` - (Optional) Type: `string`.
* `multihop` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_bgptpl`.
* `nexthop_choice` - (Optional) Type: `string`.
* `no_client_to_client_reflection` - (Optional) Type: `string`.
* `no_early_cut` - (Optional) Type: `string`.
* `output_affinity` - (Optional) Type: `string`.
* `output_filter` - (Optional) Type: `string`.
* `output_network` - (Optional) Type: `string`.
* `output_redistribute` - (Optional) Type: `string`.
* `output_selection_policy` - (Optional) Type: `string`.
* `remove_private_as` - (Optional) Type: `string`.
* `router_id` - (Optional) Type: `string`. Default: `1.1.1.1`.
* `routing_table` - (Optional) Type: `string`.
* `templates` - (Optional) Type: `string`.
* `use_bfd` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `default` - Type: `bool`.
* `invalid` - Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_bgp_template.example '*3'

# Named router
terraform import routeros_routing_bgp_template.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_bgp_template.example 'home/my-resource-name'
```
