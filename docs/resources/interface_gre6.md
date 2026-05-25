---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_gre6"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_gre6

Manages the RouterOS `/interface/gre6` menu.

## Example Usage

```terraform
resource "routeros_interface_gre6" "gre6_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # ipsec_secret = "REDACTED"
  # local_address = "10.99.0.1"
  # mtu = "replace-me"
  # name = "example"
  # remote_address = "10.99.0.1"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `ipsec_secret` - (Optional) Type: `string`.
* `local_address` - (Optional) Type: `string`.
* `mtu` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `remote_address` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_gre6.example '*3'

# Named router
terraform import routeros_interface_gre6.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_gre6.example 'home/my-resource-name'
```
