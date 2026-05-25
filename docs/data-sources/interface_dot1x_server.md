---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_dot1x_server"
description: |-
  802.1X server attaches to a specific Ethernet interface; values vary per device. Skipped.
---

# Data Source: routeros_interface_dot1x_server

802.1X server attaches to a specific Ethernet interface; values vary per device. Skipped.

## Example Usage

```terraform
data "routeros_interface_dot1x_server" "server_example" {
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
* `accounting` - (Optional) Type: `bool`. Default: `1`.
* `auth_timeout` - (Optional) Type: `string`. Default: `6000`.
* `auth_types` - (Optional) Type: `string`. Default: `1`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `guest_vlan_id` - (Optional) Type: `string`.
* `interface` - (Optional) Type: `string`.
* `interim_update` - (Optional) Type: `duration`.
* `mac_auth_mode` - (Optional) Type: `enum(mac as username|mac as username and password)`.
* `radius_mac_format` - (Optional) Type: `enum(XX:XX:XX:XX:XX:XX|XX-XX-XX-XX-XX-XX|XXXXXXXXXXXX|xx:xx:xx:xx:xx:xx|xx-xx-xx-xx-xx-xx|xxxxxxxxxxxx)`.
* `reauth_timeout` - (Optional) Type: `string`.
* `reject_vlan_id` - (Optional) Type: `string`.
* `retrans_timeout` - (Optional) Type: `string`. Default: `3000`.
* `server_fail_vlan_id` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

