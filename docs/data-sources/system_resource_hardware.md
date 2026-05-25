---
subcategory: "System"
page_title: "RouterOS: routeros_system_resource_hardware"
description: |-
  Discovered read-only menu
---

# Data Source: routeros_system_resource_hardware

Discovered read-only menu

## Example Usage

```terraform
data "routeros_system_resource_hardware" "hardware_example" {
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

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

