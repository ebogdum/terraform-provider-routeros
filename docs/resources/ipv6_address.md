---
subcategory: "IPv6"
page_title: "RouterOS: routeros_ipv6_address"
description: |-
  RouterOS resource.
---

# Resource: routeros_ipv6_address

Manages the RouterOS `/ipv6/address` menu.

## Example Usage

```terraform
resource "routeros_ipv6_address" "address_example" {
  # router = "my-router"  # which router to target; omit for the default
  address = "fd00:db8::1/64"
  interface = "ether1"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # advertise = true
  # auto_link_local = false
  # dynglob = "replace-me"
  # eui_64 = false
  # from_pool = "replace-me"
  # no_dad = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Required) Type: `cidr`. Default: `fd00:db8::1/64`.
* `advertise` - (Optional) Type: `bool`. Default: `1`.
* `auto_link_local` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `dynglob` - (Optional) Type: `string`.
* `eui_64` - (Optional) Type: `bool`.
* `from_pool` - (Optional) Type: `string`.
* `interface` - (Required) Type: `string`.
* `no_dad` - (Optional) Type: `bool`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `actual_interface` - Type: `string`.
* `deprecated` - Type: `bool`.
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.
* `link_local` - Type: `bool`.
* `preferred` - Type: `string`.
* `scope` - Type: `int`.
* `slave` - Type: `bool`.
* `valid` - Type: `string`.
* `vrf` - Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ipv6_address.example '*3'

# Named router
terraform import routeros_ipv6_address.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ipv6_address.example 'home/my-resource-name'
```
