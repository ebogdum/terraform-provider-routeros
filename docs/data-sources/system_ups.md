---
subcategory: "System"
page_title: "RouterOS: routeros_system_ups"
description: |-
  RouterOS resource.
---

# Data Source: routeros_system_ups

Manages the RouterOS `/system/ups` menu.

## Example Usage

```terraform
data "routeros_system_ups" "ups_example" {
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
* `alarm_setting` - (Optional) Type: `enum(immediate|delayed|low battery|none)`. UPS sound alarm setting: delayed - alarm is delayed to the on-battery event immediate - alarm immediately after the on-battery event low-battery - alarm only when the battery is low none - do not alarm.
* `check_capabilities` - (Optional) Type: `string`. Whether to check UPS capabilities before reading information. Disabling it can fix compatibility issues with some UPS models. (Applies to RouterOS version 6, implemented since v6.17).
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `min_runtime` - (Optional) Type: `string`. Minimal run time remaining. After a 'utility' failure, the router will monitor the runtime-left value. When the value reaches the min-runtime value, the router will go to hibernate mode. never - the router will go to hibernate mode when the "battery low" signal is sent indicating that the battery power is below 10% 0s - the router will continue to work as long as the battery is supplying sufficient voltage.
* `name` - (Optional) Type: `string`.
* `offline_time` - (Optional) Type: `string`. How long to work on batteries. The router waits that amount of time and then goes into hibernate mode until the UPS reports that the 'utility' power is back 0s - the router will go into hibernate mode according to the min-runtime setting. In this case, the router will wait until the UPS reports that the battery power is below 10%.
* `port` - (Optional) Type: `string`. Communication port of the router.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

