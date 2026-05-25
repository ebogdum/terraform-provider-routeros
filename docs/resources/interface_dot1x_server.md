---
page_title: "RouterOS: routeros_interface_dot1x_server"
description: |-
  802.1X server attaches to a specific Ethernet interface; values vary per device. Skipped.
---

# Resource: routeros_interface_dot1x_server

802.1X server attaches to a specific Ethernet interface; values vary per device. Skipped.

## Example Usage

```terraform
resource "routeros_interface_dot1x_server" "server_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # accounting = true
  # auth_timeout = "6000"
  # auth_types = "1"
  # guest_vlan_id = "replace-me"
  # interface = "ether1"
  # interim_update = "1h"
  # mac_auth_mode = "mac as username"
  # radius_mac_format = "XX:XX:XX:XX:XX:XX"
  # reauth_timeout = "replace-me"
  # reject_vlan_id = "replace-me"
  # retrans_timeout = "3000"
  # server_fail_vlan_id = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
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

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_dot1x_server.example '*3'

# Named router
terraform import routeros_interface_dot1x_server.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_dot1x_server.example 'home/my-resource-name'
```
