---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_proxy_connections"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_proxy_connections

Manages the RouterOS `/ip/proxy/connections` menu.

## Example Usage

```terraform
data "routeros_ip_proxy_connections" "connections_example" {
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
* `client` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dst_address` - (Optional) Type: `string`.
* `last_protocol` - (Optional) Type: `enum(|http/1.0|http/1.1|ftp)`.
* `rx_bytes` - (Optional) Type: `string`.
* `server` - (Optional) Type: `bool`.
* `src_address` - (Optional) Type: `string`.
* `state` - (Optional) Type: `enum(rx-header|resolving|connecting|waiting|rx-body|tx-header, ...)`.
* `tx_bytes` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

