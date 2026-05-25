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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `alarm_setting` - (Optional) Type: `enum(immediate|delayed|low-battery|none)`. UPS sound alarm setting: delayed - alarm is delayed to the on-battery event immediate - alarm immediately after the on-battery event low-battery - alarm only when the battery is low none - do not alarm.
* `beep` - (Optional) Type: `string`.
* `check_capabilities` - (Optional) Type: `string`. Whether to check UPS capabilities before reading information. Disabling it can fix compatibility issues with some UPS models. (Applies to RouterOS version 6, implemented since v6.17).
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `min_runtime` - (Optional) Type: `string`. Minimal run time remaining. After a 'utility' failure, the router will monitor the runtime-left value. When the value reaches the min-runtime value, the router will go to hibernate mode. never - the router will go to hibernate mode when the "battery low" signal is sent indicating that the battery power is below 10% 0s - the router will continue to work as long as the battery is supplying sufficient voltage.
* `name` - (Optional) Type: `string`.
* `offline_time` - (Optional) Type: `string`. How long to work on batteries. The router waits that amount of time and then goes into hibernate mode until the UPS reports that the 'utility' power is back 0s - the router will go into hibernate mode according to the min-runtime setting. In this case, the router will wait until the UPS reports that the battery power is below 10%.
* `port` - (Optional) Type: `string`. Communication port of the router.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `battery_charge` - Type: `int`.
* `battery_voltage` - Type: `string`.
* `frequency` - Type: `int`.
* `invalid` - Type: `bool`.
* `line_voltage` - Type: `string`.
* `load` - Type: `int`.
* `low_battery` - Type: `bool`.
* `manufacture_date` - Type: `string`.
* `model` - Type: `string`.
* `nominal_battery_voltage` - Type: `int`.
* `offline_after` - Type: `duration`.
* `on_battery` - Type: `bool`.
* `on_line` - Type: `bool`.
* `ouput_voltage` - Type: `string`.
* `overload` - Type: `bool`.
* `replace_battery` - Type: `bool`.
* `run_time_left` - Type: `duration`.
* `serial_number` - Type: `string`.
* `smart_boost` - Type: `bool`.
* `smart_trim` - Type: `bool`.
* `temperature` - Type: `string`.
* `transfer_cause` - Type: `string`.
* `version` - Type: `string`.

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
