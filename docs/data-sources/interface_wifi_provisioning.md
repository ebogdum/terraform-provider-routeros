---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_wifi_provisioning"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_wifi_provisioning

Manages the RouterOS `/interface/wifi/provisioning` menu.

## Example Usage

```terraform
data "routeros_interface_wifi_provisioning" "provisioning_example" {
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
* `action` - (Required) Type: `enum(none|create enabled|create disabled|create dynamic enabled)`. Default: `create-dynamic-enabled`.
* `address_ranges` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `common_name_regexp` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `identity_regexp` - (Optional) Type: `string`.
* `master_configuration` - (Optional) Type: `string`.
* `multi_link_mode` - (Optional) Type: `string`.
* `name_format` - (Optional) Type: `string`.
* `radio_mac` - (Optional) Type: `string`.
* `slave_configurations` - (Optional) Type: `string`.
* `slave_name_format` - (Optional) Type: `string`.
* `supported_bands` - (Optional) Type: `string`.
* `supported_hw_caps` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

