---
page_title: "RouterOS: routeros_ppp_profile"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ppp_profile

Manages the RouterOS `/ppp/profile` menu.

## Example Usage

```terraform
data "routeros_ppp_profile" "profile_example" {
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

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `default` - Type: `bool`.

