---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_wifi_radio"
description: |-
  Radios are hardware-backed; can't be added in software.
---

# Data Source: routeros_interface_wifi_radio

Radios are hardware-backed; can't be added in software.

## Example Usage

```terraform
data "routeros_interface_wifi_radio" "radio_example" {
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
* `x2g_channels` - (Optional) Type: `string`.
* `x5g_channels` - (Optional) Type: `string`.
* `x6g_channels` - (Optional) Type: `string`.
* `bands` - (Optional) Type: `string`.
* `cap` - (Optional) Type: `string`.
* `ciphers` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `copy_to_provisioning` - (Optional) Type: `string`.
* `countries` - (Optional) Type: `string`.
* `current_channels` - (Optional) Type: `string`.
* `current_country` - (Optional) Type: `string`.
* `current_gopclasses` - (Optional) Type: `string`.
* `current_max_reg_power` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `hw_caps` - (Optional) Type: `string`.
* `hw_type` - (Optional) Type: `string`.
* `interface` - (Optional) Type: `string`.
* `local` - (Optional) Type: `bool`.
* `max_interfaces` - (Optional) Type: `int`.
* `max_peers` - (Optional) Type: `int`.
* `max_station_interfaces` - (Optional) Type: `int`.
* `max_vlans` - (Optional) Type: `int`.
* `min_antenna_gain` - (Optional) Type: `int`.
* `phy_id` - (Optional) Type: `int`.
* `provision` - (Optional) Type: `string`.
* `radio_mac` - (Optional) Type: `string`.
* `rx_chains` - (Optional) Type: `string`.
* `tx_chains` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

