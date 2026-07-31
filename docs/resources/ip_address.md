---
subcategory: "IP"
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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `actual_interface` - (Read-only) Type: `string`.
* `address` - (Required) Type: `string`.
* `broadcast` - (Optional) Type: `string`. RouterOS `broadcast`.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `dynamic` - (Read-only) Type: `bool`.
* `interface` - (Required) Type: `string`.
* `invalid` - (Read-only) Type: `bool`.
* `netmask` - (Optional) Type: `string`. RouterOS `netmask`.
* `network` - (Optional) Type: `string`.
* `slave` - (Read-only) Type: `bool`.
* `vrf` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
