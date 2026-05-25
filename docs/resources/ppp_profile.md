---
subcategory: "PPP"
page_title: "RouterOS: routeros_ppp_profile"
description: |-
  RouterOS resource.
---

# Resource: routeros_ppp_profile

Manages the RouterOS `/ppp/profile` menu.

## Example Usage

```terraform
resource "routeros_ppp_profile" "profile_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "example"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # address_list = "replace-me"
  # bridge = "bridge1"
  # bridge_horizon = "replace-me"
  # bridge_learning = "no"
  # bridge_path_cost = "replace-me"
  # bridge_port_priority = "replace-me"
  # bridge_port_trusted = "replace-me"
  # bridge_port_vid = "replace-me"
  # change_tcp_mss = "no"
  # dhcpv6_lease_time = "replace-me"
  # dhcpv6_pd_pool = "replace-me"
  # dhcpv6_use_radius = "replace-me"
  # dns_server = "1.1.1.1,8.8.8.8"
  # idle_timeout = "replace-me"
  # incoming_filter = "replace-me"
  # insert_queue_before = "replace-me"
  # interface_list = "replace-me"
  # local_address = "10.99.0.1"
  # on_down = "replace-me"
  # on_up = "replace-me"
  # only_one = "no"
  # outgoing_filter = "replace-me"
  # parent_queue = "replace-me"
  # remote_address = "10.99.0.1"
  # remote_ipv6_prefix_pool = "replace-me"
  # remote_ipv6_prefix_reuse = "replace-me"
  # session_timeout = "replace-me"
  # use_compression = "no"
  # use_encryption = "no"
  # use_ipv6 = "yes"
  # use_mpls = "no"
  # use_upnp = "replace-me"
  # wins_server = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address_list` - (Optional) Type: `string`.
* `bridge` - (Optional) Type: `string`.
* `bridge_horizon` - (Optional) Type: `string`.
* `bridge_learning` - (Optional) Type: `enum(no|yes|default)`. Default: `4.294967295e+09`.
* `bridge_path_cost` - (Optional) Type: `string`.
* `bridge_port_priority` - (Optional) Type: `string`.
* `bridge_port_trusted` - (Optional) Type: `string`.
* `bridge_port_vid` - (Optional) Type: `string`.
* `change_tcp_mss` - (Optional) Type: `enum(no|yes|default)`. Default: `4.294967295e+09`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dhcpv6_lease_time` - (Optional) Type: `string`.
* `dhcpv6_pd_pool` - (Optional) Type: `string`.
* `dhcpv6_use_radius` - (Optional) Type: `string`.
* `dns_server` - (Optional) Type: `string`.
* `idle_timeout` - (Optional) Type: `string`.
* `incoming_filter` - (Optional) Type: `string`.
* `insert_queue_before` - (Optional) Type: `string`.
* `interface_list` - (Optional) Type: `string`.
* `local_address` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_pppprof`.
* `on_down` - (Optional) Type: `string`.
* `on_up` - (Optional) Type: `string`.
* `only_one` - (Optional) Type: `enum(no|yes|default)`. Default: `4.294967295e+09`.
* `outgoing_filter` - (Optional) Type: `string`.
* `parent_queue` - (Optional) Type: `string`.
* `remote_address` - (Optional) Type: `string`.
* `remote_ipv6_prefix_pool` - (Optional) Type: `string`.
* `remote_ipv6_prefix_reuse` - (Optional) Type: `string`.
* `session_timeout` - (Optional) Type: `string`.
* `use_compression` - (Optional) Type: `enum(no|yes|default)`. Default: `4.294967295e+09`.
* `use_encryption` - (Optional) Type: `enum(no|yes|required|default)`. Default: `4.294967295e+09`.
* `use_ipv6` - (Optional) Type: `enum(no|yes|required|default)`. Default: `1`.
* `use_mpls` - (Optional) Type: `enum(no|yes|required|default)`. Default: `4.294967295e+09`.
* `use_upnp` - (Optional) Type: `string`.
* `wins_server` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `default` - Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ppp_profile.example '*3'

# Named router
terraform import routeros_ppp_profile.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ppp_profile.example 'home/my-resource-name'
```
