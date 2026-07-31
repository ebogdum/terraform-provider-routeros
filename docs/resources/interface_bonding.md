---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_bonding"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Resource: routeros_interface_bonding

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
resource "routeros_interface_bonding" "bonding_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # mode = "replace-me"
  # mtu = "replace-me"
  # name = "tf-example"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `arp` - (Optional) Type: `string`.
* `arp_interval` - (Optional) Type: `string`. RouterOS `arp-interval`.
* `arp_ip_targets` - (Optional) Type: `string`. RouterOS `arp-ip-targets`.
* `arp_timeout` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `down_delay` - (Optional) Type: `string`. RouterOS `down-delay`.
* `forced_mac_address` - (Optional) Type: `string`. RouterOS `forced-mac-address`.
* `lacp_mode` - (Optional) Type: `string`. RouterOS `lacp-mode`.
* `lacp_rate` - (Optional) Type: `string`. RouterOS `lacp-rate`.
* `lacp_system_id` - (Optional) Type: `string`. RouterOS `lacp-system-id`.
* `lacp_system_priority` - (Optional) Type: `string`. RouterOS `lacp-system-priority`.
* `lacp_user_key` - (Optional) Type: `string`. RouterOS `lacp-user-key`.
* `link_monitoring` - (Optional) Type: `string`. RouterOS `link-monitoring`.
* `mii_interval` - (Optional) Type: `string`. RouterOS `mii-interval`.
* `min_links` - (Optional) Type: `string`. RouterOS `min-links`.
* `mlag_id` - (Optional) Type: `string`. RouterOS `mlag-id`.
* `mode` - (Optional) Type: `string`.
* `mtu` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `primary` - (Optional) Type: `string`. RouterOS `primary`.
* `slaves` - (Optional) Type: `string`. RouterOS `slaves`.
* `transmit_hash_policy` - (Optional) Type: `string`. RouterOS `transmit-hash-policy`.
* `up_delay` - (Optional) Type: `string`. RouterOS `up-delay`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_bonding.example '*3'

# Named router
terraform import routeros_interface_bonding.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_bonding.example 'home/my-resource-name'
```
