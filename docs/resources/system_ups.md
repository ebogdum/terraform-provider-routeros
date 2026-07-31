---
subcategory: "System"
page_title: "RouterOS: routeros_system_ups"
description: |-
  RouterOS resource.
---

# Resource: routeros_system_ups

Manages the RouterOS `/system/ups` menu.

## Example Usage

```terraform
resource "routeros_system_ups" "ups_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # alarm_setting = "immediate"
  # beep = "replace-me"
  # check_capabilities = "replace-me"
  # min_runtime = "replace-me"
  # name = "tf-example"
  # offline_time = "replace-me"
  # port = "443"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `alarm_setting` - (Optional) Type: `string`. UPS sound alarm setting: delayed - alarm is delayed to the on-battery event immediate - alarm immediately after the on-battery event low-battery - alarm only when the battery is low none - do not alarm
* `battery_charge` - (Read-only) Type: `int`.
* `battery_voltage` - (Read-only) Type: `string`.
* `beep` - (Optional) Type: `string`.
* `check_capabilities` - (Optional) Type: `string`. Whether to check UPS capabilities before reading information. Disabling it can fix compatibility issues with some UPS models. (Applies to RouterOS version 6, implemented since v6.17)
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `frequency` - (Read-only) Type: `int`.
* `invalid` - (Read-only) Type: `bool`.
* `line_voltage` - (Read-only) Type: `string`.
* `load` - (Read-only) Type: `int`.
* `low_battery` - (Read-only) Type: `bool`.
* `manufacture_date` - (Read-only) Type: `string`.
* `min_runtime` - (Optional) Type: `string`. Minimal run time remaining. After a 'utility' failure, the router will monitor the runtime-left value. When the value reaches the min-runtime value, the router will go to hibernate mode. never - the router will go to hibernate mode when the "battery low" signal is sent indicating that the battery power is below 10% 0s - the router will continue to work as long as the battery is supplying sufficient voltage
* `model` - (Read-only) Type: `string`.
* `name` - (Optional) Type: `string`.
* `nominal_battery_voltage` - (Read-only) Type: `int`.
* `offline_after` - (Read-only) Type: `string`.
* `offline_time` - (Optional) Type: `string`. How long to work on batteries. The router waits that amount of time and then goes into hibernate mode until the UPS reports that the 'utility' power is back 0s - the router will go into hibernate mode according to the min-runtime setting. In this case, the router will wait until the UPS reports that the battery power is below 10%
* `on_battery` - (Read-only) Type: `bool`.
* `on_line` - (Read-only) Type: `bool`.
* `ouput_voltage` - (Read-only) Type: `string`.
* `overload` - (Read-only) Type: `bool`.
* `port` - (Optional) Type: `string`. Communication port of the router.
* `replace_battery` - (Read-only) Type: `bool`.
* `run_time_left` - (Read-only) Type: `string`.
* `serial_number` - (Read-only) Type: `string`.
* `smart_boost` - (Read-only) Type: `bool`.
* `smart_trim` - (Read-only) Type: `bool`.
* `temperature` - (Read-only) Type: `string`.
* `transfer_cause` - (Read-only) Type: `string`.
* `version` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_system_ups.example '*3'

# Named router
terraform import routeros_system_ups.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_system_ups.example 'home/my-resource-name'
```
