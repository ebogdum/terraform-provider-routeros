---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_igmp_proxy_interface"
description: |-
  RouterOS resource.
---

# Data Source: routeros_routing_igmp_proxy_interface

Manages the RouterOS `/routing/igmp-proxy/interface` menu.

## Example Usage

```terraform
data "routeros_routing_igmp_proxy_interface" "interface_example" {
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
* `alternative_subnets` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `interface` - (Optional) Type: `string`.
* `threshold` - (Optional) Type: `int`. Default: `1`.
* `upstream` - (Optional) Type: `bool`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

