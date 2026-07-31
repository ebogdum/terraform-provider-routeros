---
subcategory: "IPv6"
page_title: "RouterOS: routeros_ipv6_pool"
description: |-
  RouterOS resource.
---

# Resource: routeros_ipv6_pool

Manages the RouterOS `/ipv6/pool` menu.

## Example Usage

```terraform
resource "routeros_ipv6_pool" "pool_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"
  prefix = "fd00:db8::/56"
  prefix_length = 64

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # from_pool = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `actual_prefix` - (Read-only) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dynamic` - (Read-only) Type: `bool`.
* `from_pool` - (Optional) Type: `string`.
* `invalid` - (Read-only) Type: `bool`.
* `name` - (Required) Type: `string`.
* `preferred_lifetime` - (Read-only) Type: `string`.
* `prefix` - (Required) Type: `string`.
* `prefix_length` - (Required) Type: `int`.
* `valid_lifetime` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ipv6_pool.example '*3'

# Named router
terraform import routeros_ipv6_pool.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ipv6_pool.example 'home/my-resource-name'
```
