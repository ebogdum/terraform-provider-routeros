---
page_title: "RouterOS: routeros_ip_hotspot_ip_binding"
description: |-
  Hotspot ip-binding requires existing hotspot.
---

# Data Source: routeros_ip_hotspot_ip_binding

Hotspot ip-binding requires existing hotspot.

## Example Usage

```terraform
data "routeros_ip_hotspot_ip_binding" "ip_binding_example" {
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
* `address` - (Optional) Type: `cidr`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `mac_address` - (Optional) Type: `string`.
* `server` - (Optional) Type: `string`.
* `to_address` - (Optional) Type: `ip`.
* `type` - (Optional) Type: `enum(regular|bypassed|blocked)`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

