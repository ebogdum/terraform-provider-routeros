---
subcategory: "MPLS"
page_title: "RouterOS: routeros_mpls_ldp_interface"
description: |-
  RouterOS resource.
---

# Resource: routeros_mpls_ldp_interface

Manages the RouterOS `/mpls/ldp/interface` menu.

## Example Usage

```terraform
resource "routeros_mpls_ldp_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"

  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Required) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_mpls_ldp_interface.example '*3'

# Named router
terraform import routeros_mpls_ldp_interface.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_mpls_ldp_interface.example 'home/my-resource-name'
```
