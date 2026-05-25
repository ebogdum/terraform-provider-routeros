---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_wifi"
description: |-
  WiFi virtual interface needs exactly one of radio-mac or master-interface. Skipped -- requires WiFi-enabled hardware.
---

# Data Source: routeros_interface_wifi

WiFi virtual interface needs exactly one of radio-mac or master-interface. Skipped -- requires WiFi-enabled hardware.

## Example Usage

```terraform
data "routeros_interface_wifi" "wifi_example" {
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
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `name` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

