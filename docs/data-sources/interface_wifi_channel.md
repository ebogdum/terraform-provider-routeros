---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_channel"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_wifi_channel

Manages the RouterOS `/interface/wifi/channel` menu.

## Example Usage

```terraform
data "routeros_interface_wifi_channel" "channel_example" {
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
* `band` - (Optional) Type: `enum(2ghz-b|2ghz-only-g|2ghz-b/g|5ghz-a|5ghz-only-n|5ghz-a/n|2ghz-only-n|2ghz-b/g/n|5ghz-a/n/ac|5ghz-only-ac|2ghz-g/n|5ghz-n/ac|2ghz-g|2ghz-n|2ghz-ax|2ghz-be|5ghz-n|5ghz-ac|5ghz-an|5ghz-ax|5ghz-be|6ghz-ax|6ghz-be)`. Default: `7`.
* `channel_width` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `deprioritize_unii_3_4` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `frequency` - (Optional) Type: `string`. Default: `2.412e+06`.
* `name` - (Required) Type: `string`. Default: `tf_acc_wchan`.
* `reselect_interval` - (Optional) Type: `string`.
* `reselect_time` - (Optional) Type: `string`.
* `secondary_frequency` - (Optional) Type: `string`.
* `skip_dfs_channels` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

