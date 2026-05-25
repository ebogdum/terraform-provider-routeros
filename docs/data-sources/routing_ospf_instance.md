---
subcategory: "OSPF"
page_title: "RouterOS: routeros_routing_ospf_instance"
description: |-
  RouterOS resource.
---

# Data Source: routeros_routing_ospf_instance

Manages the RouterOS `/routing/ospf/instance` menu.

## Example Usage

```terraform
data "routeros_routing_ospf_instance" "instance_example" {
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
* `domain_id` - (Optional) Type: `string`.
* `domain_tag` - (Optional) Type: `string`.
* `in_filter` - (Optional) Type: `string`.
* `mpls_te_address` - (Optional) Type: `string`.
* `mpls_te_area` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `originate_default` - (Optional) Type: `string`.
* `out_filter` - (Optional) Type: `string`.
* `out_filter_select` - (Optional) Type: `string`.
* `redistribute` - (Optional) Type: `string`.
* `router_id` - (Optional) Type: `string`.
* `routing_table` - (Optional) Type: `string`.
* `version` - (Optional) Type: `enum(2|3)`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `invalid` - Type: `bool`.

