---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_dot1x_client"
description: |-
  802.1X client attaches to a specific Ethernet interface; values vary per device. Skipped.
---

# Data Source: routeros_interface_dot1x_client

802.1X client attaches to a specific Ethernet interface; values vary per device. Skipped.

## Example Usage

```terraform
data "routeros_interface_dot1x_client" "client_example" {
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
* `anon_identity` - (Optional) Type: `string`.
* `certificate` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `eap_methods` - (Optional) Type: `string`.
* `identity` - (Optional) Type: `string`.
* `interface` - (Optional) Type: `string`.
* `password` - (Optional) Type: `string`. **Sensitive.**

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `invalid` - Type: `bool`.

