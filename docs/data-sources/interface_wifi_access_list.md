---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_access_list"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_wifi_access_list

Manages the RouterOS `/interface/wifi/access-list` menu.

## Example Usage

```terraform
data "routeros_interface_wifi_access_list" "access_list_example" {
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
* `action` - (Optional) Type: `string`.
* `allow_signal_out_of_range` - (Optional) Type: `string`.
* `client_isolation` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Optional) Type: `string`.
* `mac_address` - (Optional) Type: `string`.
* `mac_address_mask` - (Optional) Type: `string`.
* `multi_passphrase_group` - (Optional) Type: `string`.
* `passphrase` - (Optional) Type: `string`. **Sensitive.**
* `radius_accounting` - (Optional) Type: `string`.
* `signal_range` - (Optional) Type: `string`.
* `ssid_regexp` - (Optional) Type: `string`.
* `time` - (Optional) Type: `string`.
* `vlan_id` - (Optional) Type: `string`.
* `weekdays` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`. **Marked sensitive**: this menu holds a secret, which RouterOS returns in the row like any other column, so an unprojected read puts it in your state file. Use `proplist` to name the columns you need.
* `last_logged_in` - Type: `string`.
* `last_logged_out` - Type: `string`.
* `times_matched` - Type: `string`.

