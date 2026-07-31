---
subcategory: "System & misc"
page_title: "RouterOS: routeros_lcd"
description: |-
  Requires LCD-equipped board (e.g. RB1100AHx4)
---

# Resource: routeros_lcd

Requires LCD-equipped board (e.g. RB1100AHx4)

## Example Usage

```terraform
resource "routeros_lcd" "lcd_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # backlight_timeout = "replace-me"
  # color_scheme = "replace-me"
  # default_screen = "replace-me"
  # enabled = "replace-me"
  # read_only_mode = "replace-me"
  # time_interval = "replace-me"
  # touch_screen = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `backlight_timeout` - (Optional) Type: `string`. Time after which LCD touchscreen is turned off
* `color_scheme` - (Optional) Type: `string`. Changes to color scheme with a dark or light background.
* `default_screen` - (Optional) Type: `string`. Default screen that is showed after startup.
* `enabled` - (Optional) Type: `string`. Turns LCD touchscreen on/off. When off, it stops and resets statistics gathering and closes the LCD program.
* `read_only_mode` - (Optional) Type: `string`. Enables or disables Read-Only mode. If Read-Only mode is enabled, then menus which can be used to change configuration are hidden.
* `time_interval` - (Optional) Type: `string`. Time interval of displayed interface statistics in Stats screen
* `touch_screen` - (Optional) Type: `string`. Enable/disable touch screen input.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_lcd.example '*3'

# Named router
terraform import routeros_lcd.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_lcd.example 'home/my-resource-name'
```
