---
page_title: "RouterOS: routeros_system_device_mode"
description: |-
  RouterOS resource.
---

# Data Source: routeros_system_device_mode

Manages the RouterOS `/system/device-mode` menu.

## Example Usage

```terraform
data "routeros_system_device_mode" "device_mode_example" {
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
* `allowed_versions` - (Optional) Type: `list`.
* `attempt_count` - (Optional) Type: `int`.
* `bandwidth_test` - (Optional) Type: `bool`.
* `container` - (Optional) Type: `bool`.
* `email` - (Optional) Type: `bool`.
* `fetch` - (Optional) Type: `bool`.
* `flagged` - (Optional) Type: `bool`.
* `flagging_enabled` - (Optional) Type: `bool`.
* `hotspot` - (Optional) Type: `bool`.
* `install_any_version` - (Optional) Type: `bool`.
* `ipsec` - (Optional) Type: `bool`.
* `l2tp` - (Optional) Type: `bool`.
* `mode` - (Optional) Type: `string`.
* `partitions` - (Optional) Type: `bool`.
* `pptp` - (Optional) Type: `bool`.
* `proxy` - (Optional) Type: `bool`.
* `romon` - (Optional) Type: `bool`.
* `routerboard` - (Optional) Type: `bool`.
* `scheduler` - (Optional) Type: `bool`.
* `smb` - (Optional) Type: `bool`.
* `sniffer` - (Optional) Type: `bool`.
* `socks` - (Optional) Type: `bool`.
* `traffic_gen` - (Optional) Type: `bool`.
* `zerotier` - (Optional) Type: `bool`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

