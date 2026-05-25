---
subcategory: "MPLS"
page_title: "RouterOS: routeros_mpls_forwarding_table"
description: |-
  RouterOS resource.
---

# Data Source: routeros_mpls_forwarding_table

Manages the RouterOS `/mpls/forwarding-table` menu.

## Example Usage

```terraform
data "routeros_mpls_forwarding_table" "forwarding_table_example" {
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
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `label` - (Optional) Type: `int`.
* `ldp` - (Optional) Type: `bool`.
* `nexthops` - (Optional) Type: `string`.
* `prefix` - (Optional) Type: `string`.
* `te_sender` - (Optional) Type: `string`.
* `te_session` - (Optional) Type: `string`.
* `traffic_eng` - (Optional) Type: `bool`.
* `type` - (Optional) Type: `enum(|LDP|VPN|Traffic Eng.|VPLS)`.
* `vpls` - (Optional) Type: `string`.
* `vpn` - (Optional) Type: `bool`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

