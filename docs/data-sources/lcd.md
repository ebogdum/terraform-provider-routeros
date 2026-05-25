---
page_title: "RouterOS: routeros_lcd"
description: |-
  Requires LCD-equipped board (e.g. RB1100AHx4)
---

# Data Source: routeros_lcd

Requires LCD-equipped board (e.g. RB1100AHx4)

## Example Usage

```terraform
data "routeros_lcd" "lcd_example" {
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
* `backlight_timeout` - (Optional) Type: `string`. Time after which LCD touchscreen is turned off.
* `color_scheme` - (Optional) Type: `string`. Changes to color scheme with a dark or light background.
* `default_screen` - (Optional) Type: `string`. Default screen that is showed after startup.
* `enabled` - (Optional) Type: `string`. Turns LCD touchscreen on/off. When off, it stops and resets statistics gathering and closes the LCD program.
* `read_only_mode` - (Optional) Type: `string`. Enables or disables Read-Only mode. If Read-Only mode is enabled, then menus which can be used to change configuration are hidden.
* `time_interval` - (Optional) Type: `string`. Time interval of displayed interface statistics in Stats screen.
* `touch_screen` - (Optional) Type: `string`. Enable/disable touch screen input.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

