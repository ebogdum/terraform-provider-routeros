---
subcategory: "Firewall"
page_title: "RouterOS: routeros_ip_firewall_connection"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_firewall_connection

Manages the RouterOS `/ip/firewall/connection` menu.

## Example Usage

```terraform
data "routeros_ip_firewall_connection" "connection_example" {
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
* `assured` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `confirmed` - (Optional) Type: `bool`.
* `connection_mark` - (Optional) Type: `string`.
* `connection_type` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dst_address` - (Optional) Type: `ip`.
* `dst_port` - (Optional) Type: `int`.
* `dstnat` - (Optional) Type: `bool`.
* `dying` - (Optional) Type: `bool`.
* `expected` - (Optional) Type: `bool`.
* `fasttrack` - (Optional) Type: `bool`.
* `helper_used` - (Optional) Type: `bool`.
* `hw_offload` - (Optional) Type: `bool`.
* `orig_repl_bytes` - (Optional) Type: `string`.
* `orig_repl_fasttrack_bytes` - (Optional) Type: `string`.
* `orig_repl_fasttrack_packets` - (Optional) Type: `string`.
* `orig_repl_packets` - (Optional) Type: `string`.
* `orig_repl_rate` - (Optional) Type: `string`.
* `protocol` - (Optional) Type: `int`.
* `reply_dst_address` - (Optional) Type: `ip`.
* `reply_dst_port` - (Optional) Type: `int`.
* `reply_src_address` - (Optional) Type: `ip`.
* `reply_src_port` - (Optional) Type: `int`.
* `seen_reply` - (Optional) Type: `bool`.
* `src_address` - (Optional) Type: `ip`.
* `src_port` - (Optional) Type: `int`.
* `srcnat` - (Optional) Type: `bool`.
* `timeout` - (Optional) Type: `duration`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

