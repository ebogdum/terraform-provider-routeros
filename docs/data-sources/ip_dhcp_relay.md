---
subcategory: "DHCP"
page_title: "RouterOS: routeros_ip_dhcp_relay"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_dhcp_relay

Manages the RouterOS `/ip/dhcp-relay` menu.

## Example Usage

```terraform
data "routeros_ip_dhcp_relay" "dhcp_relay_example" {
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
* `add_relay_info` - (Optional) Type: `bool`.
* `delay_threshold` - (Optional) Type: `duration`.
* `dhcp_server` - (Required) Type: `string`. Default: `127.0.0.1`.
* `dhcp_server_vrf` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Required) Type: `string`.
* `local_address` - (Optional) Type: `ip`.
* `name` - (Required) Type: `string`. Default: `tf-acc-relay`.
* `relay_info_remote_id` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

