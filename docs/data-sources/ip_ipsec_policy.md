---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_ipsec_policy"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Data Source: routeros_ip_ipsec_policy

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
data "routeros_ip_ipsec_policy" "policy_example" {
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
* `action` - (Optional) Type: `enum(discard|none|encrypt)`. Default: `2`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `dst_address` - (Optional) Type: `cidr`.
* `dst_port` - (Optional) Type: `int`.
* `group` - (Optional) Type: `string`.
* `ipsec_protocols` - (Optional) Type: `enum(|ah|esp)`. Default: `2`.
* `level` - (Optional) Type: `enum(use|require|unique)`. Default: `2`.
* `peer` - (Optional) Type: `string`.
* `proposal` - (Optional) Type: `string`.
* `protocol` - (Optional) Type: `enum(icmp|igmp|ggp|ip-encap|tcp|egp, ...)`. Default: `255`.
* `src_address` - (Optional) Type: `cidr`.
* `src_port` - (Optional) Type: `int`.
* `template` - (Optional) Type: `bool`.
* `tunnel` - (Optional) Type: `bool`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `default` - Type: `bool`.
* `dynamic` - Type: `bool`.
* `sa_dst_address` - Type: `string`.
* `sa_src_address` - Type: `string`.

