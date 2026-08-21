---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_security_multi_passphrase"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_wifi_security_multi_passphrase

Manages the RouterOS `/interface/wifi/security/multi-passphrase` menu.

Each row is one passphrase entry in a PPSK group. Assign the group name to `multi_passphrase_group`
on `routeros_interface_wifi`, `routeros_interface_wifi_configuration`, `routeros_interface_wifi_security`
or `routeros_interface_wifi_access_list` to put it to use.

## Example Usage

```terraform
resource "routeros_interface_wifi_security_multi_passphrase" "multi_passphrase_example" {
  # router = "my-router"  # which router to target; omit for the default
  group      = "guest-group"
  passphrase = "replace-me"

  # Optional attributes (uncomment as needed):
  # comment = "managed by terraform"
  # disabled = false
  # expires = "replace-me"
  # isolation = false
  # vlan_id = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `expires` - (Optional) Type: `string`. Expiration date and time for this passphrase; doesn't affect the rest of the group. Once reached, existing clients using it are disconnected and new clients can't use it. Leave unset for no expiry.
* `group` - (Required) Type: `string`. The PPSK group name. Assigning it to a security profile or an access list enables use of every passphrase defined under it.
* `isolation` - (Optional) Type: `bool`. Whether a client using this passphrase is isolated from other clients on the AP: traffic from an isolated client is not forwarded to other clients, and unicast traffic from a non-isolated client is not forwarded to an isolated one.
* `passphrase` - (Optional) Type: `string`. **Sensitive.** The PSK passphrase. Multiple entries may share a passphrase. Not compatible with WPA3-PSK.
* `vlan_id` - (Optional) Type: `string`. VLAN ID assigned to clients using this passphrase. Only supported on wifi-qcom interfaces; a wifi-qcom-ac AP will refuse a client whose passphrase carries one.

## Attribute Reference

* `id` - RouterOS internal .id.
* `expired` - (Read-only) Type: `bool`. Whether this passphrase's `expires` date/time has passed.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_wifi_security_multi_passphrase.example '*3'

# Named router
terraform import routeros_interface_wifi_security_multi_passphrase.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_wifi_security_multi_passphrase.example 'home/my-resource-name'
```
