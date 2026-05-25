---
page_title: "RouterOS: routeros_routing_nexthop"
description: |-
  RouterOS resource.
---

# Data Source: routeros_routing_nexthop

Manages the RouterOS `/routing/nexthop` menu.

## Example Usage

```terraform
data "routeros_routing_nexthop" "nexthop_example" {
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
* `address` - (Optional) Type: `string`.
* `afi` - (Optional) Type: `string`.
* `bgp_vpn` - (Optional) Type: `bool`.
* `check_gateway` - (Optional) Type: `enum(none|arp|ping|bfd)`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `flap_count` - (Optional) Type: `string`.
* `gateway_check_ok` - (Optional) Type: `bool`.
* `gw_address` - (Optional) Type: `string`.
* `gw_check_ok` - (Optional) Type: `bool`.
* `immediate_gw_address` - (Optional) Type: `string`.
* `immediate_gw_blackhole` - (Optional) Type: `bool`.
* `immediate_gw_flap_count` - (Optional) Type: `int`.
* `immediate_gw_interface_idx` - (Optional) Type: `int`.
* `immediate_gw_mpls_label` - (Optional) Type: `int`.
* `immediate_gw_mpls_peer_id` - (Optional) Type: `int`.
* `immediate_gw_weight` - (Optional) Type: `int`.
* `interface_ok` - (Optional) Type: `bool`.
* `mpls_label` - (Optional) Type: `string`.
* `mpls_peer_id` - (Optional) Type: `string`.
* `reachable` - (Optional) Type: `bool`.
* `scope` - (Optional) Type: `int`.
* `target_scope` - (Optional) Type: `int`.
* `unresolved` - (Optional) Type: `bool`.
* `weight` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

