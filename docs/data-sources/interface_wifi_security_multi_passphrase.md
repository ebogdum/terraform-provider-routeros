---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_security_multi_passphrase"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_wifi_security_multi_passphrase

Manages the RouterOS `/interface/wifi/security/multi-passphrase` menu.

## Example Usage

```terraform
data "routeros_interface_wifi_security_multi_passphrase" "multi_passphrase_example" {
  # router = "my-router"  # omit for the default router
  # filter = { group = "guest-group" }
}
```

## Argument Reference

This data source supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to query.
* `filter` - (Optional) Map of field=value pairs to narrow the result set.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `expires` - (Optional) Type: `string`.
* `group` - (Optional) Type: `string`.
* `isolation` - (Optional) Type: `bool`.
* `passphrase` - (Optional) Type: `string`. **Sensitive.**
* `vlan_id` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `expired` - Type: `bool`.
