---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_kid_control_device"
description: |-
  Discovered; needs kid-control-name fixture
---

# Data Source: routeros_ip_kid_control_device

Discovered; needs kid-control-name fixture

## Example Usage

```terraform
data "routeros_ip_kid_control_device" "device_example" {
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
* `blocked` - (Optional) Type: `bool`.
* `bytes` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `mac_address` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `rate_limited` - (Optional) Type: `bool`.
* `rate_up_down` - (Optional) Type: `string`.
* `reset_counters` - (Optional) Type: `string`.
* `user` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `activity` - Type: `string`.
* `dynamic` - Type: `bool`.
* `ip_address` - Type: `string`.

