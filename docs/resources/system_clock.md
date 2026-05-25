---
page_title: "RouterOS: routeros_system_clock"
description: |-
  Setting clock outside automated test scope -- would skew router time.
---

# Resource: routeros_system_clock

Setting clock outside automated test scope -- would skew router time.

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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `date` - (Optional) Type: `string`.
* `time` - (Optional) Type: `string`.
* `time_zone_autodetect` - (Optional) Type: `bool`.
* `time_zone_name` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `dst_active` - Type: `bool`.
* `gmt_offset` - Type: `string`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_clock.this 'home'
```
