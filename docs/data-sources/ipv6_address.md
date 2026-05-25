---
subcategory: "IPv6"
page_title: "RouterOS: routeros_ipv6_address"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ipv6_address

Manages the RouterOS `/ipv6/address` menu.

## Example Usage

```terraform
data "routeros_ipv6_address" "address_example" {
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
* `address` - (Required) Type: `cidr`. Default: `fd00:db8::1/64`.
* `advertise` - (Optional) Type: `bool`. Default: `1`.
* `auto_link_local` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `eui_64` - (Optional) Type: `bool`.
* `from_pool` - (Optional) Type: `string`.
* `interface` - (Required) Type: `string`.
* `no_dad` - (Optional) Type: `bool`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `actual_interface` - Type: `string`.
* `deprecated` - Type: `bool`.
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.
* `link_local` - Type: `bool`.
* `slave` - Type: `bool`.
* `vrf` - Type: `string`.

