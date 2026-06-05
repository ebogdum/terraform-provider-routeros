---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_channel"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_wifi_channel

Manages the RouterOS `/interface/wifi/channel` menu.

## Example Usage

```terraform
resource "routeros_interface_wifi_channel" "channel_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # band = "5ghz-ax"
  # channel_width = "replace-me"
  # deprioritize_unii_3_4 = "replace-me"
  # frequency = "2.412e+06"
  # reselect_interval = "replace-me"
  # reselect_time = "replace-me"
  # secondary_frequency = "replace-me"
  # skip_dfs_channels = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
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

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_wifi_channel.example '*3'

# Named router
terraform import routeros_interface_wifi_channel.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_wifi_channel.example 'home/my-resource-name'
```
