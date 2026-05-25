---
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
  # frame_types = "admit all"
  # horizon = 0
  # hw = "replace-me"
  # ingress_filtering = true
  # interface = "ether1"
  # internal_path_cost = "replace-me"
  # learn = "auto"
  # multicast_router = "Disabled"
  # mvrp_applicant_state = "normal participant"
  # mvrp_registrar_state = "normal"
  # path_cost = "replace-me"
  # point_to_point = "auto"
  # priority = 128
  # pvid = 0
  # restricted_role = false
  # restricted_tcn = false
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
* `edge` - (Optional) Type: `enum(auto|yes|no|yes discover|no discover)`.
* `fast_leave` - (Optional) Type: `bool`.
* `frame_types` - (Optional) Type: `enum(admit all|admit only VLAN tagged|admit only untagged and priority tagged)`.
* `horizon` - (Optional) Type: `int`.
* `hw` - (Optional) Type: `string`.
* `ingress_filtering` - (Optional) Type: `bool`. Default: `1`.
* `interface` - (Optional) Type: `string`.
* `internal_path_cost` - (Optional) Type: `string`.
* `learn` - (Optional) Type: `enum(auto|no|yes)`.
* `multicast_router` - (Optional) Type: `enum(Disabled|Temporary Query|Permanent)`.
* `mvrp_applicant_state` - (Optional) Type: `enum(normal participant|non participant)`.
* `mvrp_registrar_state` - (Optional) Type: `enum(normal|fixed)`.
* `path_cost` - (Optional) Type: `string`.
* `point_to_point` - (Optional) Type: `enum(auto|yes|no)`.
* `priority` - (Optional) Type: `int`. Default: `128`.
* `pvid` - (Optional) Type: `int`.
* `restricted_role` - (Optional) Type: `bool`.
* `restricted_tcn` - (Optional) Type: `bool`.
* `tag_stacking` - (Optional) Type: `bool`.
* `trusted` - (Optional) Type: `bool`.
* `trusted_ra` - (Optional) Type: `bool`.
* `unknown_multicast_flood` - (Optional) Type: `bool`. Default: `1`.
* `unknown_unicast_flood` - (Optional) Type: `bool`. Default: `1`.

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
