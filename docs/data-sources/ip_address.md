---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_address"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_address

Manages the RouterOS `/ip/address` menu.

## Example Usage

```terraform
data "routeros_ip_address" "address_example" {
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
* `address` - (Required) Type: `cidr`.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `interface` - (Required) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `actual_interface` - Type: `string`.
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.
* `network` - Type: `ip`.
* `slave` - Type: `bool`.
* `vrf` - Type: `string`.

