---
subcategory: "Firewall"
page_title: "RouterOS: routeros_ipv6_firewall_address_list"
description: |-
  RouterOS resource.
---

# Resource: routeros_ipv6_firewall_address_list

Manages the RouterOS `/ipv6/firewall/address-list` menu.

## Example Usage

```terraform
resource "routeros_ipv6_firewall_address_list" "address_list_example" {
  # router = "my-router"  # which router to target; omit for the default
  address = "fd00:db8::/64"
  list = "my-list"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # parent = 0
  # timeout = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address` - (Required) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `creation_time` - (Read-only) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dynamic` - (Read-only) Type: `bool`.
* `list` - (Required) Type: `string`.
* `parent` - (Read-only) Type: `int`.
* `timeout` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ipv6_firewall_address_list.example '*3'

# Named router
terraform import routeros_ipv6_firewall_address_list.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ipv6_firewall_address_list.example 'home/my-resource-name'
```
