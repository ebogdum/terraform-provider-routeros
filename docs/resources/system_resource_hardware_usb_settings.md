---
subcategory: "System"
page_title: "RouterOS: routeros_system_resource_hardware_usb_settings"
description: |-
  Mirrors RouterOS /system/resource/hardware/usb-settings.
---

# Resource: routeros_system_resource_hardware_usb_settings

Mirrors RouterOS `/system/resource/hardware/usb-settings`.

## Example Usage

```terraform
resource "routeros_system_resource_hardware_usb_settings" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # authorization = true
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `authorization` - (Optional) Type: `bool`. RouterOS `authorization`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_resource_hardware_usb_settings.this 'home'
```
