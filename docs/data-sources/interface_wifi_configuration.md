---
page_title: "RouterOS: routeros_interface_wifi_configuration"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_wifi_configuration

Manages the RouterOS `/interface/wifi/configuration` menu.

## Example Usage

```terraform
data "routeros_interface_wifi_configuration" "configuration_example" {
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
* `antenna_gain` - (Optional) Type: `string`.
* `channel` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `country` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `distance` - (Optional) Type: `string`.
* `mode` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_wcfg`.
* `ssid` - (Optional) Type: `string`.
* `tx_power` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

