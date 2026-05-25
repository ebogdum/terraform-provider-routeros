---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_configuration"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_wifi_configuration

Manages the RouterOS `/interface/wifi/configuration` menu.

## Example Usage

```terraform
resource "routeros_interface_wifi_configuration" "configuration_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # antenna_gain = "replace-me"
  # channel = "replace-me"
  # country = "replace-me"
  # distance = "replace-me"
  # mode = "replace-me"
  # ssid = "replace-me"
  # tx_power = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
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

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_wifi_configuration.example '*3'

# Named router
terraform import routeros_interface_wifi_configuration.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_wifi_configuration.example 'home/my-resource-name'
```
