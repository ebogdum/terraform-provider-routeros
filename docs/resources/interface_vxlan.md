---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_vxlan"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_vxlan

Manages the RouterOS `/interface/vxlan` menu.

## Example Usage

```terraform
resource "routeros_interface_vxlan" "vxlan_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "example"
  vni = "100"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # bridge = "bridge1"
  # interface = "ether1"
  # local_address = "10.99.0.1"
  # mac_address = "10.99.0.0/24"
  # mtu = "replace-me"
  # port = "443"
  # ttl = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `arp` - (Optional) Type: `string`.
* `arp_timeout` - (Optional) Type: `string`.
* `bridge` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Optional) Type: `string`.
* `local_address` - (Optional) Type: `string`.
* `mac_address` - (Optional) Type: `string`.
* `mtu` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_vxlan`.
* `port` - (Optional) Type: `string`.
* `ttl` - (Optional) Type: `string`.
* `vni` - (Required) Type: `string`. Default: `100`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_vxlan.example '*3'

# Named router
terraform import routeros_interface_vxlan.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_vxlan.example 'home/my-resource-name'
```
