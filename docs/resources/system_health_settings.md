---
subcategory: "System"
page_title: "RouterOS: routeros_system_health_settings"
description: |-
  Mirrors RouterOS /system/health/settings.
---

# Resource: routeros_system_health_settings

Mirrors RouterOS `/system/health/settings`.

## Example Usage

```terraform
resource "routeros_system_health_settings" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # cpu_overtemp_check = true
  # cpu_overtemp_startup_delay = "replace-me"
  # cpu_overtemp_threshold = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `cpu_overtemp_check` - (Optional) Type: `bool`. RouterOS `cpu-overtemp-check`.
* `cpu_overtemp_startup_delay` - (Optional) Type: `string`. RouterOS `cpu-overtemp-startup-delay`.
* `cpu_overtemp_threshold` - (Optional) Type: `string`. RouterOS `cpu-overtemp-threshold`.
* `fan_control_interval` - (Optional) Type: `string`. RouterOS `fan-control-interval`.
* `fan_full_speed_temp` - (Optional) Type: `string`. RouterOS `fan-full-speed-temp`.
* `fan_min_speed_percent` - (Optional) Type: `string`. RouterOS `fan-min-speed-percent`.
* `fan_mode` - (Optional) Type: `string`. RouterOS `fan-mode`.
* `fan_on_threshold` - (Optional) Type: `string`. RouterOS `fan-on-threshold`.
* `fan_switch` - (Optional) Type: `string`. RouterOS `fan-switch`.
* `fan_target_temp` - (Optional) Type: `string`. RouterOS `fan-target-temp`.
* `use_fan` - (Optional) Type: `string`. RouterOS `use-fan`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_health_settings.this 'home'
```
