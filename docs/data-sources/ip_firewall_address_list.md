---
page_title: "RouterOS: routeros_ip_firewall_address_list"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_firewall_address_list

Manages the RouterOS `/ip/firewall/address-list` menu.

## Example Usage

```terraform
data "routeros_ip_firewall_address_list" "address_list_example" {
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
* `address` - (Required) Type: `string`. Default: `10.255.255.0/30`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `list` - (Required) Type: `string`. Default: `tf_acc_list`.
* `timeout` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `dynamic` - Type: `bool`.

