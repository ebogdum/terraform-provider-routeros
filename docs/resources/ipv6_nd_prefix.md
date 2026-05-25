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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `x6to4_interface` - (Optional) Type: `string`. Default: `4.294967295e+09`.
* `autonomous` - (Optional) Type: `bool`. Default: `1`.
* `dhcpv6_pd_preferred` - (Optional) Type: `bool`.
* `disabled` - (Optional) Type: `bool`.
* `interface` - (Optional) Type: `string`.
* `no6to4` - (Optional) Type: `string`.
* `on_link` - (Optional) Type: `bool`. Default: `1`.
* `preferred_lifetime` - (Optional) Type: `duration`. Default: `604800`.
* `prefix` - (Optional) Type: `string`.
* `valid_lifetime` - (Optional) Type: `duration`. Default: `2.592e+06`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.

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
