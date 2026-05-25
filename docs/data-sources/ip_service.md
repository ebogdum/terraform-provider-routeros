---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_service"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_service

Manages the RouterOS `/ip/service` menu.

## Example Usage

```terraform
data "routeros_ip_service" "service_example" {
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
* `address` - (Optional) Type: `string`.
* `available_from` - (Optional) Type: `string`.
* `cert` - (Optional) Type: `string`.
* `certificate` - (Optional) Type: `string`. Default: `4.294967295e+09`.
* `conn` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `max_sessions` - (Optional) Type: `int`.
* `nondyn` - (Optional) Type: `string`.
* `port` - (Optional) Type: `int`.
* `tls_version` - (Optional) Type: `enum(any|only-v1.2)`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `connection` - Type: `bool`.
* `container` - Type: `string`.
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.
* `local` - Type: `ip`.
* `name` - Type: `string`.
* `net_ns` - Type: `int`.
* `proto` - Type: `string`.
* `protocol` - Type: `enum(tcp|udp)`.
* `remote` - Type: `string`.

