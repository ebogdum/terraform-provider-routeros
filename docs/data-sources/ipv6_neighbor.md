---
subcategory: "IPv6"
page_title: "RouterOS: routeros_ipv6_neighbor"
description: |-
  IPv6 neighbor table — read-only on most devices.
---

# Data Source: routeros_ipv6_neighbor

IPv6 neighbor table — read-only on most devices.

## Example Usage

```terraform
data "routeros_ipv6_neighbor" "neighbor_example" {
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
* `address` - (Optional) Type: `ipv6`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Optional) Type: `string`.
* `mac_address` - (Optional) Type: `string`.
* `mac_ping` - (Optional) Type: `string`.
* `mac_telnet` - (Optional) Type: `string`.
* `make_static` - (Optional) Type: `string`.
* `ping` - (Optional) Type: `string`.
* `router` - (Optional) Type: `bool`.
* `telnet` - (Optional) Type: `string`.
* `torch` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `bridge_port` - Type: `string`.
* `dynamic` - Type: `bool`.
* `host_name` - Type: `string`.
* `status` - Type: `string`.
* `vrf` - Type: `string`.

