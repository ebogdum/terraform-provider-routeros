---
subcategory: "MPLS"
page_title: "RouterOS: routeros_mpls_interface"
description: |-
  RouterOS resource.
---

# Resource: routeros_mpls_interface

Manages the RouterOS `/mpls/interface` menu.

## Example Usage

```terraform
resource "routeros_mpls_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # builtin = false
  # input = "replace-me"
  # mpls_mtu = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `builtin` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `input` - (Optional) Type: `string`.
* `interface` - (Required) Type: `string`.
* `mpls_mtu` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_mpls_interface.example '*3'

# Named router
terraform import routeros_mpls_interface.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_mpls_interface.example 'home/my-resource-name'
```
