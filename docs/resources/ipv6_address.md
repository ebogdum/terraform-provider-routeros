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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `actual_interface` - (Read-only) Type: `string`.
* `address` - (Required) Type: `string`.
* `advertise` - (Optional) Type: `bool`.
* `auto_link_local` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `deprecated` - (Read-only) Type: `bool`.
* `disabled` - (Optional) Type: `bool`.
* `dynamic` - (Read-only) Type: `bool`.
* `dynglob` - (Read-only) Type: `string`.
* `eui_64` - (Optional) Type: `bool`.
* `from_pool` - (Optional) Type: `string`.
* `from_pool_policy` - (Optional) Type: `string`. RouterOS `from-pool-policy`.
* `interface` - (Required) Type: `string`.
* `invalid` - (Read-only) Type: `bool`.
* `link_local` - (Read-only) Type: `bool`.
* `no_dad` - (Optional) Type: `bool`.
* `preferred` - (Read-only) Type: `string`.
* `scope` - (Read-only) Type: `int`.
* `slave` - (Read-only) Type: `bool`.
* `valid` - (Read-only) Type: `string`.
* `vrf` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
