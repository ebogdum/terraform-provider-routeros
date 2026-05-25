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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `arp` - (Optional) Type: `enum(disabled|enabled|proxy-arp|reply-only|local-proxy-arp)`. Default: `1`.
* `arp_timeout` - (Optional) Type: `duration`.
* `bgp_signaled` - (Optional) Type: `bool`.
* `bridge` - (Optional) Type: `string`.
* `bridge_cost` - (Optional) Type: `string`.
* `bridge_horizon` - (Optional) Type: `string`.
* `bridge_pvid` - (Optional) Type: `string`.
* `cisco_bgp_signaled` - (Optional) Type: `bool`.
* `cisco_static_id` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `mac_address` - (Optional) Type: `string`.
* `mtu` - (Optional) Type: `int`. Default: `1500`.
* `pw_control_word` - (Optional) Type: `string`.
* `pw_l2mtu` - (Optional) Type: `string`.
* `pw_type` - (Optional) Type: `string`.
* `remote_peer` - (Optional) Type: `string`.
* `vpls_id` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `bgp_vpls` - Type: `string`.
* `bgp_vpls_prefix` - Type: `string`.
* `local_label` - Type: `int`.
* `remote_group` - Type: `int`.
* `remote_label` - Type: `int`.
* `remote_status` - Type: `string`.
* `te_tunnel` - Type: `int`.

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
