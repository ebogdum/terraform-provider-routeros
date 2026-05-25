---
subcategory: "IPv6"
page_title: "RouterOS: routeros_ipv6_pool"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ipv6_pool

Manages the RouterOS `/ipv6/pool` menu.

## Example Usage

```terraform
data "routeros_ipv6_pool" "pool_example" {
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
* `from_pool` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_pool6`.
* `prefix` - (Required) Type: `string`. Default: `fd00:db8::/56`.
* `prefix_length` - (Required) Type: `int`. Default: `64`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `actual_prefix` - Type: `string`.
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.
* `preferred_lifetime` - Type: `string`.
* `valid_lifetime` - Type: `string`.

