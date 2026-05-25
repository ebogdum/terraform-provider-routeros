---
page_title: "RouterOS: routeros_routing_bgp_template"
description: |-
  RouterOS resource.
---

# Data Source: routeros_routing_bgp_template

Manages the RouterOS `/routing/bgp/template` menu.

## Example Usage

```terraform
data "routeros_routing_bgp_template" "template_example" {
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
* `as` - (Optional) Type: `string`. Default: `65000`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `hold_time` - (Optional) Type: `string`.
* `keepalive_time` - (Optional) Type: `string`.
* `multihop` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_bgptpl`.
* `nexthop_choice` - (Optional) Type: `string`.
* `router_id` - (Optional) Type: `string`. Default: `1.1.1.1`.
* `routing_table` - (Optional) Type: `string`.
* `templates` - (Optional) Type: `string`.
* `use_bfd` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `default` - Type: `bool`.

