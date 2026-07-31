---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_kid_control_device"
description: |-
  Discovered; needs kid-control-name fixture
---

# Resource: routeros_ip_kid_control_device

Discovered; needs kid-control-name fixture

## Example Usage

```terraform
resource "routeros_ip_kid_control_device" "device_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # blocked = false
  # bytes = "replace-me"
  # mac_address = "10.99.0.0/24"
  # name = "tf-example"
  # rate_limited = false
  # rate_up_down = "replace-me"
  # reset_counters = "replace-me"
  # user = "myuser"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `activity` - (Read-only) Type: `string`.
* `blocked` - (Read-only) Type: `bool`.
* `bytes` - (Read-only) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `dynamic` - (Read-only) Type: `bool`.
* `ip_address` - (Read-only) Type: `string`.
* `mac_address` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `rate_limited` - (Read-only) Type: `bool`.
* `rate_up_down` - (Read-only) Type: `string`.
* `reset_counters` - (Read-only) Type: `string`.
* `user` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_kid_control_device.example '*3'

# Named router
terraform import routeros_ip_kid_control_device.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_kid_control_device.example 'home/my-resource-name'
```
