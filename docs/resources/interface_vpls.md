---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_vpls"
description: |-
  VPLS needs MPLS/LDP setup; skipped from automated acc tests.
---

# Resource: routeros_interface_vpls

VPLS needs MPLS/LDP setup; skipped from automated acc tests.

## Example Usage

```terraform
resource "routeros_interface_vpls" "vpls_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "enabled"
  # arp_timeout = "1h"
  # bgp_signaled = false
  # bridge = "bridge1"
  # bridge_cost = "replace-me"
  # bridge_horizon = "replace-me"
  # bridge_pvid = "replace-me"
  # cisco_bgp_signaled = false
  # cisco_static_id = "replace-me"
  # mac_address = "10.99.0.0/24"
  # mtu = 1500
  # pw_control_word = "replace-me"
  # pw_l2mtu = "replace-me"
  # pw_type = "replace-me"
  # remote_peer = "replace-me"
  # vpls_id = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `arp` - (Optional) Type: `string`.
* `arp_timeout` - (Optional) Type: `string`.
* `bgp_signaled` - (Read-only) Type: `bool`.
* `bgp_vpls` - (Read-only) Type: `string`.
* `bgp_vpls_prefix` - (Read-only) Type: `string`.
* `bridge` - (Optional) Type: `string`.
* `bridge_cost` - (Optional) Type: `string`.
* `bridge_horizon` - (Optional) Type: `string`.
* `bridge_pvid` - (Optional) Type: `string`.
* `cisco_bgp_signaled` - (Read-only) Type: `bool`.
* `cisco_static_id` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disable_running_check` - (Optional) Type: `string`. RouterOS `disable-running-check`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `local_label` - (Read-only) Type: `int`.
* `mac_address` - (Optional) Type: `string`.
* `mtu` - (Optional) Type: `int`.
* `name` - (Optional) Type: `string`. RouterOS `name`.
* `peer` - (Optional) Type: `string`. RouterOS `peer`.
* `pw_control_word` - (Optional) Type: `string`.
* `pw_l2mtu` - (Optional) Type: `string`.
* `pw_type` - (Optional) Type: `string`.
* `remote_group` - (Read-only) Type: `int`.
* `remote_label` - (Read-only) Type: `int`.
* `remote_peer` - (Read-only) Type: `string`.
* `remote_status` - (Read-only) Type: `string`.
* `te_tunnel` - (Read-only) Type: `int`.
* `vpls_id` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_vpls.example '*3'

# Named router
terraform import routeros_interface_vpls.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_vpls.example 'home/my-resource-name'
```
