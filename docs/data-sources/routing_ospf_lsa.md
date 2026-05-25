---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_ospf_lsa"
description: |-
  RouterOS resource.
---

# Data Source: routeros_routing_ospf_lsa

Manages the RouterOS `/routing/ospf/lsa` menu.

## Example Usage

```terraform
data "routeros_routing_ospf_lsa" "lsa_example" {
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
* `age` - (Optional) Type: `int`.
* `area` - (Optional) Type: `string`.
* `body` - (Optional) Type: `string`.
* `checksum` - (Optional) Type: `int`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `flushing` - (Optional) Type: `bool`.
* `instance` - (Optional) Type: `string`.
* `link` - (Optional) Type: `string`.
* `link_instance_id` - (Optional) Type: `int`.
* `originator` - (Optional) Type: `ip`.
* `self_originated` - (Optional) Type: `bool`.
* `sequence` - (Optional) Type: `int`.
* `type` - (Optional) Type: `string`.
* `wraparound` - (Optional) Type: `bool`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `dynamic` - Type: `bool`.

