---
subcategory: "System"
page_title: "RouterOS: routeros_system_clock"
description: |-
  Setting clock outside automated test scope — would skew router time.
---

# Resource: routeros_system_clock

Setting clock outside automated test scope — would skew router time.

## Example Usage

```terraform
resource "routeros_system_clock" "clock_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # date = "replace-me"
  # time = "replace-me"
  # time_zone_autodetect = false
  # time_zone_name = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `date` - (Optional) Type: `string`.
* `dst_active` - (Optional) Type: `bool`.
* `gmt_offset` - (Optional) Type: `string`.
* `time` - (Optional) Type: `string`.
* `time_zone_autodetect` - (Optional) Type: `bool`.
* `time_zone_name` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_clock.this 'home'
```
