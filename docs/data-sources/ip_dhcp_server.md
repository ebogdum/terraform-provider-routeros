---
subcategory: "DHCP"
page_title: "RouterOS: routeros_ip_dhcp_server"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_dhcp_server

Manages the RouterOS `/ip/dhcp-server` menu.

## Example Usage

```terraform
data "routeros_ip_dhcp_server" "dhcp_server_example" {
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
* `address_pool` - (Optional) Type: `string`.
* `allow_dual_stack_queue` - (Optional) Type: `bool`. Default: `1`.
* `always_broadcast` - (Optional) Type: `bool`.
* `authoritative` - (Optional) Type: `enum(yes|after 2s delay|after 10s delay|no)`. Default: `0`.
* `bootp_lease_time` - (Optional) Type: `duration`. Default: `4.294967295e+09`.
* `bootp_support` - (Optional) Type: `enum(none|static|dynamic)`. Default: `1`.
* `client_mac_limit` - (Optional) Type: `int`. Default: `4.294967295e+09`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `conflict_detection` - (Optional) Type: `bool`. Default: `1`.
* `delay_threshold` - (Optional) Type: `duration`.
* `dhcp_option_set` - (Optional) Type: `string`. Default: `4.294967295e+09`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dynamic_lease_identifiers` - (Optional) Type: `string`.
* `insert_queue_before` - (Optional) Type: `string`. Default: `0`.
* `interface` - (Required) Type: `string`.
* `lease_script` - (Optional) Type: `string`.
* `lease_time` - (Optional) Type: `duration`. Default: `1800`.
* `name` - (Required) Type: `string`. Default: `tf_acc_dhcps`.
* `parent_queue` - (Optional) Type: `string`.
* `relay` - (Optional) Type: `ip`.
* `server_address` - (Optional) Type: `ip`.
* `use_framed_as_classless` - (Optional) Type: `bool`. Default: `1`.
* `use_radius` - (Optional) Type: `enum(no|yes|accounting)`.
* `use_reconfigure` - (Optional) Type: `bool`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

