---
page_title: "RouterOS: routeros_interface_l2tp_ether"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_l2tp_ether

Manages the RouterOS `/interface/l2tp-ether` menu.

## Example Usage

```terraform
resource "routeros_interface_l2tp_ether" "l2tp_ether_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # ipsec_secret = "REDACTED"
  # local_address = "10.99.0.1"
  # mac_address = "10.99.0.0/24"
  # mtu = "replace-me"
  # name = "tf-example"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `ipsec_secret` - (Optional) Type: `string`.
* `local_address` - (Optional) Type: `string`.
* `mac_address` - (Optional) Type: `string`.
* `mtu` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_l2tp_ether.example '*3'

# Named router
terraform import routeros_interface_l2tp_ether.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_l2tp_ether.example 'home/my-resource-name'
```
