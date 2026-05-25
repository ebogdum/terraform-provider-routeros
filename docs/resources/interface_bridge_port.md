---
subcategory: "Bridge"
page_title: "RouterOS: routeros_interface_bridge_port"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Resource: routeros_interface_bridge_port

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
resource "routeros_interface_bridge_port" "port_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # auto_isolate = false
  # bpdu_guard = false
  # bridge = "bridge1"
  # broadcast_flood = true
  # edge = "auto"
  # fast_leave = false
  # frame_types = "admit-all"
  # hardware_offload = true
  # horizon = 0
  # hw = "replace-me"
  # inactive = false
  # ingress_filtering = true
  # interface = "ether1"
  # internal_path_cost = "replace-me"
  # learn = "auto"
  # multicast_router = "disabled"
  # mvrp_applicant_state = "normal-participant"
  # mvrp_registrar_state = "normal"
  # parent = 0
  # path_cost = "replace-me"
  # point_to_point = "auto"
  # priority = 128
  # pvid = 0
  # restricted_role = false
  # restricted_tcn = false
  # role = 0
  # status = 0
  # tag_stacking = false
  # trusted = false
  # trusted_ra = false
  # unknown_multicast_flood = true
  # unknown_unicast_flood = true
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
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

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `dynamic` - Type: `bool`.
* `hw_offload` - Type: `bool`.
* `hw_offload_group` - Type: `string`.
* `port_status` - Type: `enum(|inactive|active|disabled)`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_bridge_port.example '*3'

# Named router
terraform import routeros_interface_bridge_port.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_bridge_port.example 'home/my-resource-name'
```
