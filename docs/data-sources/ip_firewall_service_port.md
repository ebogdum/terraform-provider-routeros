---
subcategory: "Firewall"
page_title: "RouterOS: routeros_ip_firewall_service_port"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_firewall_service_port

Manages the RouterOS `/ip/firewall/service-port` menu.

## Example Usage

```terraform
data "routeros_ip_firewall_service_port" "service_port_example" {
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
* `disabled` - (Optional) Type: `bool`.
* `ports` - (Optional) Type: `int`.
* `sip` - (Optional) Type: `string`.
* `sip_direct_media` - (Optional) Type: `bool`.
* `sip_timeout` - (Optional) Type: `duration`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `invalid` - Type: `bool`.
* `name` - Type: `string`.

