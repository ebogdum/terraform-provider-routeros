---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_bgp_connection"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Data Source: routeros_routing_bgp_connection

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
data "routeros_routing_bgp_connection" "connection_example" {
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
* `afi` - (Optional) Type: `string`.
* `as` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connect` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `hold_time` - (Optional) Type: `string`.
* `instance` - (Optional) Type: `string`.
* `keepalive_time` - (Optional) Type: `string`.
* `listen` - (Optional) Type: `string`.
* `multihop` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `nexthop_choice` - (Optional) Type: `string`.
* `routing_table` - (Optional) Type: `string`.
* `tcp_md5_key` - (Optional) Type: `string`. **Sensitive.**
* `use_bfd` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

