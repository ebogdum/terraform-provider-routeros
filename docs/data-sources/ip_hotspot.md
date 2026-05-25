---
page_title: "RouterOS: routeros_ip_hotspot"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_hotspot

Manages the RouterOS `/ip/hotspot` menu.

## Example Usage

```terraform
data "routeros_ip_hotspot" "hotspot_example" {
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
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Required) Type: `string`.
* `keepalive_timeout` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Descriptive name of the profile. Default: `tf_acc_hotspot`.
* `profile` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

