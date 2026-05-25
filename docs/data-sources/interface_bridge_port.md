---
subcategory: "Bridge"
page_title: "RouterOS: routeros_interface_bridge_port"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Data Source: routeros_interface_bridge_port

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
data "routeros_interface_bridge_port" "port_example" {
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
* `auto_isolate` - (Optional) Type: `bool`.
* `bpdu_guard` - (Optional) Type: `bool`.
* `bridge` - (Optional) Type: `string`.
* `broadcast_flood` - (Optional) Type: `bool`. Default: `1`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `edge` - (Optional) Type: `enum(auto|yes|no|yes-discover|no-discover)`.
* `fast_leave` - (Optional) Type: `bool`.
* `frame_types` - (Optional) Type: `enum(admit-all|admit-only-vlan-tagged|admit-only-untagged-and-priority-tagged)`.
* `hardware_offload` - (Optional) Type: `bool`. Default: `1`.
* `horizon` - (Optional) Type: `int`.
* `hw` - (Optional) Type: `string`.
* `inactive` - (Optional) Type: `bool`.
* `ingress_filtering` - (Optional) Type: `bool`. Default: `1`.
* `interface` - (Optional) Type: `string`.
* `internal_path_cost` - (Optional) Type: `string`.
* `learn` - (Optional) Type: `enum(auto|no|yes)`.
* `multicast_router` - (Optional) Type: `enum(disabled|temporary-query|permanent)`.
* `mvrp_applicant_state` - (Optional) Type: `enum(normal-participant|non-participant)`.
* `mvrp_registrar_state` - (Optional) Type: `enum(normal|fixed)`.
* `parent` - (Optional) Type: `int`.
* `path_cost` - (Optional) Type: `string`.
* `point_to_point` - (Optional) Type: `enum(auto|yes|no)`.
* `priority` - (Optional) Type: `int`. Default: `128`.
* `pvid` - (Optional) Type: `int`.
* `restricted_role` - (Optional) Type: `bool`.
* `restricted_tcn` - (Optional) Type: `bool`.
* `role` - (Optional) Type: `int`.
* `status` - (Optional) Type: `int`.
* `tag_stacking` - (Optional) Type: `bool`.
* `trusted` - (Optional) Type: `bool`.
* `trusted_ra` - (Optional) Type: `bool`.
* `unknown_multicast_flood` - (Optional) Type: `bool`. Default: `1`.
* `unknown_unicast_flood` - (Optional) Type: `bool`. Default: `1`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `dynamic` - Type: `bool`.
* `hw_offload` - Type: `bool`.
* `hw_offload_group` - Type: `string`.
* `port_status` - Type: `enum(|inactive|active|disabled)`.

