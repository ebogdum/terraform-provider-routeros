---
subcategory: "System"
page_title: "RouterOS: routeros_system_resource_hardware"
description: |-
  Discovered read-only menu
---

# Resource: routeros_system_resource_hardware

Discovered read-only menu

## Example Usage

```terraform
resource "routeros_system_resource_hardware" "hardware_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # category = "replace-me"
  # device_id = "replace-me"
  # devices = "replace-me"
  # io = "replace-me"
  # irq = 0
  # location = "replace-me"
  # memory = "replace-me"
  # name = "example"
  # owner = "replace-me"
  # parent = 0
  # pci = "replace-me"
  # ports = 0
  # serial_number = "replace-me"
  # speed = "replace-me"
  # std_descr = "replace-me"
  # type = "USB"
  # usb = "replace-me"
  # usb_version = "replace-me"
  # vendor = "replace-me"
  # vendor_id = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `category` - (Optional) Type: `string`.
* `device_id` - (Optional) Type: `string`.
* `devices` - (Optional) Type: `string`.
* `io` - (Optional) Type: `string`.
* `irq` - (Optional) Type: `int`.
* `location` - (Optional) Type: `string`.
* `memory` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `owner` - (Optional) Type: `string`.
* `parent` - (Optional) Type: `int`. Default: `4.294967295e+09`.
* `pci` - (Optional) Type: `string`.
* `ports` - (Optional) Type: `int`.
* `serial_number` - (Optional) Type: `string`.
* `speed` - (Optional) Type: `string`.
* `std_descr` - (Optional) Type: `string`.
* `type` - (Optional) Type: `enum(USB|PCI|SCSI|Serial)`.
* `usb` - (Optional) Type: `string`.
* `usb_version` - (Optional) Type: `string`.
* `vendor` - (Optional) Type: `string`.
* `vendor_id` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_system_resource_hardware.example '*3'

# Named router
terraform import routeros_system_resource_hardware.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_system_resource_hardware.example 'home/my-resource-name'
```
