---
subcategory: "Interfaces"
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
  # mac = "replace-me"
  # mac_auth_mode = "mac-as-username"
  # radius_mac_format = "xx:xx:xx:xx:xx:xx"
  # reauth_timeout = "replace-me"
  # reject_vlan_id = "replace-me"
  # retrans_timeout = "3000"
  # server_fail_vlan_id = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `accounting` - (Optional) Type: `bool`.
* `auth_timeout` - (Optional) Type: `string`.
* `auth_types` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `guest_vlan_id` - (Optional) Type: `string`.
* `interface` - (Optional) Type: `string`.
* `interim_update` - (Optional) Type: `string`.
* `invalid` - (Read-only) Type: `bool`.
* `mac` - (Read-only) Type: `string`.
* `mac_auth_mode` - (Optional) Type: `string`.
* `radius_mac_format` - (Optional) Type: `string`.
* `reauth_timeout` - (Optional) Type: `string`.
* `reject_vlan_id` - (Optional) Type: `string`.
* `retrans_timeout` - (Optional) Type: `string`.
* `server_fail_vlan_id` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
