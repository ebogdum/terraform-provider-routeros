---
page_title: "RouterOS: routeros_system_gps"
description: |-
  Requires GPS hardware/package
---

# Data Source: routeros_system_gps

Requires GPS hardware/package

## Example Usage

```terraform
data "routeros_system_gps" "gps_example" {
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
* `channel` - (Optional) Type: `string`. Port channel used by the device.
* `coordinate_format` - (Optional) Type: `string`. Which coordinate format to use, "Decimal Degrees", "Degrees Minutes Seconds" or "NMEA format DDDMM.MM[MM]".
* `enabled` - (Optional) Type: `string`. Whether GPS is enabled.
* `gps_antenna_select` - (Optional) Type: `string`. Depending on the model. Internal antenna can be selected, if the device has one installed.
* `init_channel` - (Optional) Type: `string`. Channel for init-string execution.
* `init_string` - (Optional) Type: `string`. AT init string for GPS initialization.
* `port` - (Optional) Type: `string`. Name of the USB/Serial port where the GPS receiver is connected.
* `set_system_time` - (Optional) Type: `string`. Whether to set the router's date and time to one received from GPS.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

