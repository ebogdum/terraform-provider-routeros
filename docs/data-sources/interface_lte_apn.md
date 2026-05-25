---
page_title: "RouterOS: routeros_interface_lte_apn"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_lte_apn

Manages the RouterOS `/interface/lte/apn` menu.

## Example Usage

```terraform
data "routeros_interface_lte_apn" "apn_example" {
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
* `add_default_route` - (Optional) Type: `bool`.
* `apn` - (Required) Type: `string`. Default: `internet`.
* `authentication` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default_route_distance` - (Optional) Type: `int`.
* `ip_type` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf-acc-apn`.
* `use_network_apn` - (Optional) Type: `bool`.
* `use_peer_dns` - (Optional) Type: `bool`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `default` - Type: `bool`.

