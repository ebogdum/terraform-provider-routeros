---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_aaa"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_wifi_aaa

Manages the RouterOS `/interface/wifi/aaa` menu.

## Example Usage

```terraform
data "routeros_interface_wifi_aaa" "aaa_example" {
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
* `called_format` - (Optional) Type: `string`.
* `calling_format` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `interim_update` - (Optional) Type: `string`.
* `mac_caching` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `nas_identifier` - (Optional) Type: `string`.
* `password_format` - (Optional) Type: `string`.
* `username_format` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

