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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `accept_dynamic_neighbors` - (Optional) Type: `string`. RouterOS `accept-dynamic-neighbors`.
* `afi` - (Optional) Type: `string`. RouterOS `afi`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `hello_interval` - (Optional) Type: `string`. RouterOS `hello-interval`.
* `hold_time` - (Optional) Type: `string`. RouterOS `hold-time`.
* `interface` - (Required) Type: `string`.
* `transport_addresses` - (Optional) Type: `string`. RouterOS `transport-addresses`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
