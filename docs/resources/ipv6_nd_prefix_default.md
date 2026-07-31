---
subcategory: "IPv6"
page_title: "RouterOS: routeros_ipv6_nd_prefix_default"
description: |-
  Mirrors RouterOS /ipv6/nd/prefix/default.
---

# Resource: routeros_ipv6_nd_prefix_default

Mirrors RouterOS `/ipv6/nd/prefix/default`.

## Example Usage

```terraform
resource "routeros_ipv6_nd_prefix_default" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # autonomous = true
  # dhcp6_pd_preferred = true
  # preferred_lifetime = "replace-me"
  # valid_lifetime = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `autonomous` - (Optional) Type: `bool`. RouterOS `autonomous`.
* `dhcp6_pd_preferred` - (Optional) Type: `bool`. RouterOS `dhcp6-pd-preferred`.
* `preferred_lifetime` - (Optional) Type: `string`. RouterOS `preferred-lifetime`.
* `valid_lifetime` - (Optional) Type: `string`. RouterOS `valid-lifetime`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_ipv6_nd_prefix_default.this 'home'
```
