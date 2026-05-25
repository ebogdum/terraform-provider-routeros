---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_vrrp"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_vrrp

Manages the RouterOS `/interface/vrrp` menu.

## Example Usage

```terraform
resource "routeros_interface_vrrp" "vrrp_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  name = "example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # authentication = "replace-me"
  # interval = "replace-me"
  # password = "REDACTED"
  # priority = "replace-me"
  # remote_address = "10.99.0.1"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `arp` - (Optional) Type: `string`.
* `arp_timeout` - (Optional) Type: `string`.
* `authentication` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Required) Type: `string`.
* `interval` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_vrrp`.
* `password` - (Optional) Type: `string`.
* `priority` - (Optional) Type: `string`.
* `remote_address` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_vrrp.example '*3'

# Named router
terraform import routeros_interface_vrrp.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_vrrp.example 'home/my-resource-name'
```
