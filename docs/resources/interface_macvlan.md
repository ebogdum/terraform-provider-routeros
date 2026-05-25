---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_macvlan"
description: |-
  MACVLAN needs an existing parent interface that supports it; live values from CHR may not match. Skipped.
---

# Resource: routeros_interface_macvlan

MACVLAN needs an existing parent interface that supports it; live values from CHR may not match. Skipped.

## Example Usage

```terraform
resource "routeros_interface_macvlan" "macvlan_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # mac_address = "10.99.0.0/24"
  # mode = "replace-me"
  # mtu = "replace-me"
  # name = "example"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `arp` - (Optional) Type: `string`.
* `arp_timeout` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `mac_address` - (Optional) Type: `string`.
* `mode` - (Optional) Type: `string`.
* `mtu` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_macvlan.example '*3'

# Named router
terraform import routeros_interface_macvlan.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_macvlan.example 'home/my-resource-name'
```
