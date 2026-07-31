---
subcategory: "IPv6"
page_title: "RouterOS: routeros_ipv6_nd_prefix"
description: |-
  RouterOS resource.
---

# Resource: routeros_ipv6_nd_prefix

Manages the RouterOS `/ipv6/nd/prefix` menu.

## Example Usage

```terraform
resource "routeros_ipv6_nd_prefix" "prefix_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # x6to4_interface = "4.294967295e+09"
  # autonomous = true
  # dhcpv6_pd_preferred = false
  # interface = "ether1"
  # no6to4 = "replace-me"
  # on_link = true
  # preferred_lifetime = "604800"
  # prefix = "replace-me"
  # valid_lifetime = "2.592e+06"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `autonomous` - (Optional) Type: `bool`.
* `dhcp6_pd_preferred` - (Optional) Type: `string`. RouterOS `dhcp6-pd-preferred`.
* `dhcpv6_pd_preferred` - (Read-only) Type: `bool`.
* `disabled` - (Optional) Type: `bool`.
* `dynamic` - (Read-only) Type: `bool`.
* `interface` - (Optional) Type: `string`.
* `invalid` - (Read-only) Type: `bool`.
* `no6to4` - (Read-only) Type: `string`.
* `on_link` - (Optional) Type: `bool`.
* `preferred_lifetime` - (Optional) Type: `string`.
* `prefix` - (Optional) Type: `string`.
* `valid_lifetime` - (Optional) Type: `string`.
* `x6to4_interface` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ipv6_nd_prefix.example '*3'

# Named router
terraform import routeros_ipv6_nd_prefix.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ipv6_nd_prefix.example 'home/my-resource-name'
```
