---
subcategory: "IPv6"
page_title: "RouterOS: routeros_ipv6_dhcp_client"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ipv6_dhcp_client

Manages the RouterOS `/ipv6/dhcp-client` menu.

## Example Usage

```terraform
data "routeros_ipv6_dhcp_client" "dhcp_client_example" {
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
* `clientid` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default_route_distance` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `hostname` - (Optional) Type: `string`.
* `interface` - (Required) Type: `string`.
* `request` - (Required) Type: `string`. Default: `address`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

