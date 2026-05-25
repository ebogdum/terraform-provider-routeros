---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_steering"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_wifi_steering

Manages the RouterOS `/interface/wifi/steering` menu.

## Example Usage

```terraform
data "routeros_interface_wifi_steering" "steering_example" {
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
* `x2g_probe_delay` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `name` - (Required) Type: `string`. Default: `tf_acc_wstr`.
* `neighbor_group` - (Optional) Type: `string`.
* `neighbor_groups` - (Optional) Type: `string`.
* `rrm` - (Optional) Type: `string`.
* `transition_request_count` - (Optional) Type: `string`.
* `transition_threshold` - (Optional) Type: `string`.
* `transition_threshold_period` - (Optional) Type: `string`.
* `transition_threshold_time` - (Optional) Type: `string`.
* `transition_time` - (Optional) Type: `string`.
* `wnm` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

