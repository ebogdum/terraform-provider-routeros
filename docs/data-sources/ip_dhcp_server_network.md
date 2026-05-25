---
subcategory: "DHCP"
page_title: "RouterOS: routeros_ip_dhcp_server_network"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_dhcp_server_network

Manages the RouterOS `/ip/dhcp-server/network` menu.

## Example Usage

```terraform
data "routeros_ip_dhcp_server_network" "network_example" {
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
* `address` - (Required) Type: `cidr`. Default: `10.255.255.0/30`.
* `boot_file_name` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dhcp_option_set` - (Optional) Type: `string`. Default: `4.294967295e+09`.
* `domain` - (Optional) Type: `string`.
* `gateway` - (Optional) Type: `string`. Default: `10.255.255.1`.
* `netmask` - (Optional) Type: `string`.
* `next_server` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

