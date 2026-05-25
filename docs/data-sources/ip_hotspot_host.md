---
subcategory: "Hotspot"
page_title: "RouterOS: routeros_ip_hotspot_host"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_hotspot_host

Manages the RouterOS `/ip/hotspot/host` menu.

## Example Usage

```terraform
data "routeros_ip_hotspot_host" "host_example" {
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
* `address` - (Optional) Type: `ip`.
* `authorized` - (Optional) Type: `bool`.
* `bridge_port` - (Optional) Type: `string`.
* `bypassed` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `idle_time` - (Optional) Type: `string`.
* `mac_address` - (Optional) Type: `string`.
* `make_binding` - (Optional) Type: `string`.
* `rx_packets` - (Optional) Type: `string`.
* `rx_rate` - (Optional) Type: `string`.
* `server` - (Optional) Type: `string`.
* `to_address` - (Optional) Type: `ip`.
* `tx_packets` - (Optional) Type: `string`.
* `tx_rate` - (Optional) Type: `string`.
* `vlan_id` - (Optional) Type: `int`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `bytes_in` - Type: `string`.
* `bytes_out` - Type: `string`.
* `host_dead_time` - Type: `string`.
* `idle_timeout` - Type: `duration`.
* `keepalive_timeout` - Type: `duration`.
* `packets_in` - Type: `string`.
* `packets_out` - Type: `string`.
* `uptime` - Type: `string`.

