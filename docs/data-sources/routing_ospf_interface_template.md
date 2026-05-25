---
subcategory: "OSPF"
page_title: "RouterOS: routeros_routing_ospf_interface_template"
description: |-
  References an existing ospf area; auto-test can't synthesise.
---

# Data Source: routeros_routing_ospf_interface_template

References an existing ospf area; auto-test can't synthesise.

## Example Usage

```terraform
data "routeros_routing_ospf_interface_template" "interface_template_example" {
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
* `area` - (Optional) Type: `string`.
* `auth_id` - (Optional) Type: `string`.
* `auth_key` - (Optional) Type: `string`. **Sensitive.**
* `comment` - (Optional) Type: `string`. Free-form comment.
* `cost` - (Optional) Type: `int`.
* `dead_interval` - (Optional) Type: `duration`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `hello_interval` - (Optional) Type: `duration`.
* `instance_id` - (Optional) Type: `int`.
* `interfaces` - (Optional) Type: `string`.
* `networks` - (Optional) Type: `string`.
* `passive` - (Optional) Type: `string`.
* `prefix_list` - (Optional) Type: `string`.
* `priority` - (Optional) Type: `int`.
* `retransmit_interval` - (Optional) Type: `duration`.
* `transmit_delay` - (Optional) Type: `int`.
* `use_bfd` - (Optional) Type: `string`.
* `vlink_neighbor_id` - (Optional) Type: `string`.
* `vlink_transit_area` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

