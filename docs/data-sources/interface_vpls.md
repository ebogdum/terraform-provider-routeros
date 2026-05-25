---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_vpls"
description: |-
  VPLS needs MPLS/LDP setup; skipped from automated acc tests.
---

# Data Source: routeros_interface_vpls

VPLS needs MPLS/LDP setup; skipped from automated acc tests.

## Example Usage

```terraform
data "routeros_interface_vpls" "vpls_example" {
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

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `bgp_vpls` - Type: `string`.
* `bgp_vpls_prefix` - Type: `string`.
* `local_label` - Type: `int`.
* `remote_group` - Type: `int`.
* `remote_label` - Type: `int`.
* `remote_status` - Type: `string`.
* `te_tunnel` - Type: `int`.

