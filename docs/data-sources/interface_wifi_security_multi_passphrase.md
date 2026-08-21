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

  # Omit proplist and every PSK lands in state in cleartext.
  proplist = [".id", "group", "vlan-id"]
}
```

## Argument Reference

This data source supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to query.
* `filter` - (Optional) Map of field=value pairs to narrow the result set.
* `proplist` - (Optional) List of columns to return. Omit to return every column, `passphrase` included.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows as string maps, plus the device's `.id`. Marked sensitive: RouterOS returns
  `passphrase` in cleartext to an account holding the `sensitive` policy, so an unprojected read puts every PSK in
  your state file. Use `proplist` to leave it out:

```terraform
data "routeros_interface_wifi_security_multi_passphrase" "groups" {
  proplist = [".id", "group", "vlan-id"]
}
```
