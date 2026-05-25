---
page_title: "RouterOS: routeros_interface_wifi_access_list"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_wifi_access_list

Manages the RouterOS `/interface/wifi/access-list` menu.

## Example Usage

```terraform
resource "routeros_interface_wifi_access_list" "access_list_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "replace-me"
  # allow_signal_out_of_range = "replace-me"
  # client_isolation = "replace-me"
  # interface = "ether1"
  # mac_address = "10.99.0.0/24"
  # mac_address_mask = "replace-me"
  # multi_passphrase_group = "replace-me"
  # passphrase = "replace-me"
  # radius_accounting = "replace-me"
  # signal_range = "replace-me"
  # ssid_regexp = "replace-me"
  # time = "replace-me"
  # vlan_id = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
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

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_wifi_access_list.example '*3'

# Named router
terraform import routeros_interface_wifi_access_list.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_wifi_access_list.example 'home/my-resource-name'
```
