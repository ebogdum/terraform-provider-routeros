---
page_title: "RouterOS: routeros_ip_dhcp_client"
description: |-
  DHCP client per interface. On most devices the default config already binds one to an interface, causing "dhcp-client on that interface already" on a fresh add. Skipped.
---

# Data Source: routeros_ip_dhcp_client

DHCP client per interface. On most devices the default config already binds one to an interface, causing "dhcp-client on that interface already" on a fresh add. Skipped.

## Example Usage

```terraform
data "routeros_ip_dhcp_client" "dhcp_client_example" {
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
* `add_default_route` - (Optional) Type: `enum(no|yes|special classless)`. Default: `1`.
* `allow_reconfigure` - (Optional) Type: `bool`.
* `check_gateway` - (Optional) Type: `enum(none|arp|ping|bfd)`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default_route_distance` - (Optional) Type: `int`.
* `default_route_tables` - (Optional) Type: `string`.
* `dhcp_options` - (Optional) Type: `list`.
* `disabled` - (Optional) Type: `bool`.
* `dscp` - (Optional) Type: `int`.
* `interface` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `script` - (Optional) Type: `string`.
* `use_broadcast` - (Optional) Type: `enum(both|always|never)`.
* `use_peer_dns` - (Optional) Type: `bool`. Default: `1`.
* `use_peer_ntp` - (Optional) Type: `bool`. Default: `1`.
* `vlan_priority` - (Optional) Type: `int`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `address` - Type: `cidr`.
* `dhcp_server` - Type: `ip`.
* `dynamic` - Type: `bool`.
* `expires_after` - Type: `duration`.
* `gateway` - Type: `ip`.
* `invalid` - Type: `bool`.
* `primary_dns` - Type: `ip`.
* `primary_ntp` - Type: `ip`.
* `status` - Type: `string`.

