---
page_title: "RouterOS: routeros_ip_address"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_address

Manages the RouterOS `/ip/address` menu.

## Example Usage

```terraform
resource "routeros_ip_address" "address_example" {
  # router = "my-router"  # which router to target; omit for the default
  address = "10.99.0.0/24"
  interface = "ether1"

  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Required) Type: `cidr`.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `interface` - (Required) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `actual_interface` - Type: `string`.
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.
* `network` - Type: `ip`.
* `slave` - Type: `bool`.
* `vrf` - Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_address.example '*3'

# Named router
terraform import routeros_ip_address.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_address.example 'home/my-resource-name'
```
