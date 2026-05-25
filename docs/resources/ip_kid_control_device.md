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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `blocked` - (Optional) Type: `bool`.
* `bytes` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `mac_address` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `rate_limited` - (Optional) Type: `bool`.
* `rate_up_down` - (Optional) Type: `string`.
* `reset_counters` - (Optional) Type: `string`.
* `user` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `activity` - Type: `string`.
* `dynamic` - Type: `bool`.
* `ip_address` - Type: `string`.

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
