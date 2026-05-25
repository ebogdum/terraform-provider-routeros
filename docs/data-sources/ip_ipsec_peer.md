---
page_title: "RouterOS: routeros_ip_ipsec_peer"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_ipsec_peer

Manages the RouterOS `/ip/ipsec/peer` menu.

## Example Usage

```terraform
data "routeros_ip_ipsec_peer" "peer_example" {
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
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `exchange_mode` - (Optional) Type: `enum(base|main|aggressive|IKE2)`. Default: `2`.
* `local_address` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `passive` - (Optional) Type: `bool`.
* `port` - (Optional) Type: `int`.
* `profile` - (Optional) Type: `string`.
* `send_initial_contact` - (Optional) Type: `bool`. Default: `1`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

