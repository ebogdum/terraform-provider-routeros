---
subcategory: "System"
page_title: "RouterOS: routeros_system_gps"
description: |-
  Requires GPS hardware/package
---

# Resource: routeros_system_gps

Requires GPS hardware/package

## Example Usage

```terraform
resource "routeros_system_gps" "gps_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # channel = "replace-me"
  # coordinate_format = "replace-me"
  # enabled = "replace-me"
  # gps_antenna_select = "replace-me"
  # init_channel = "replace-me"
  # init_string = "replace-me"
  # port = "443"
  # set_system_time = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `channel` - (Optional) Type: `string`. Port channel used by the device
* `coordinate_format` - (Optional) Type: `string`. Which coordinate format to use, "Decimal Degrees", "Degrees Minutes Seconds" or "NMEA format DDDMM.MM[MM]"
* `enabled` - (Optional) Type: `string`. Whether GPS is enabled
* `gps_antenna_select` - (Optional) Type: `string`. Depending on the model. Internal antenna can be selected, if the device has one installed.
* `init_channel` - (Optional) Type: `string`. Channel for init-string execution
* `init_string` - (Optional) Type: `string`. AT init string for GPS initialization
* `port` - (Optional) Type: `string`. Name of the USB/Serial port where the GPS receiver is connected
* `set_system_time` - (Optional) Type: `string`. Whether to set the router's date and time to one received from GPS.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_system_gps.example '*3'

# Named router
terraform import routeros_system_gps.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_system_gps.example 'home/my-resource-name'
```
