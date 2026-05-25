---
subcategory: "IPv6"
page_title: "RouterOS: routeros_ipv6_dhcp_server"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ipv6_dhcp_server

Manages the RouterOS `/ipv6/dhcp-server` menu.

## Example Usage

```terraform
data "routeros_ipv6_dhcp_server" "dhcp_server_example" {
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
* `dhcp_option` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Required) Type: `string`.
* `lease_time` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_dhcp6`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

