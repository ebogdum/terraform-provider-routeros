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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `auto_isolate` - (Optional) Type: `bool`.
* `bpdu_guard` - (Optional) Type: `bool`.
* `bridge` - (Optional) Type: `string`.
* `broadcast_flood` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dynamic` - (Read-only) Type: `bool`.
* `edge` - (Optional) Type: `string`.
* `fast_leave` - (Optional) Type: `bool`.
* `frame_types` - (Optional) Type: `string`.
* `hardware_offload` - (Read-only) Type: `bool`.
* `horizon` - (Optional) Type: `string`. Split-horizon group used to isolate ports. A number, or `none` (the default).
* `hw` - (Optional) Type: `string`.
* `hw_offload` - (Read-only) Type: `bool`.
* `hw_offload_group` - (Read-only) Type: `string`.
* `inactive` - (Read-only) Type: `bool`.
* `ingress_filtering` - (Optional) Type: `bool`.
* `interface` - (Optional) Type: `string`.
* `internal_path_cost` - (Optional) Type: `string`.
* `learn` - (Optional) Type: `string`.
* `multicast_router` - (Optional) Type: `string`.
* `mvrp_applicant_state` - (Optional) Type: `string`.
* `mvrp_registrar_state` - (Optional) Type: `string`.
* `parent` - (Read-only) Type: `int`.
* `path_cost` - (Optional) Type: `string`.
* `point_to_point` - (Optional) Type: `string`.
* `port_status` - (Read-only) Type: `string`.
* `priority` - (Optional) Type: `int`.
* `pvid` - (Optional) Type: `int`.
* `restricted_role` - (Optional) Type: `bool`.
* `restricted_tcn` - (Optional) Type: `bool`.
* `role` - (Read-only) Type: `int`.
* `status` - (Read-only) Type: `int`.
* `tag_stacking` - (Optional) Type: `bool`.
* `trusted` - (Optional) Type: `bool`.
* `trusted_dhcpv6` - (Optional) Type: `string`. RouterOS `trusted-dhcpv6`.
* `trusted_ra` - (Optional) Type: `bool`.
* `unknown_multicast_flood` - (Optional) Type: `bool`.
* `unknown_unicast_flood` - (Optional) Type: `bool`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
