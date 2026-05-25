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
  name = "example"
  prefix = "fd00:db8::/56"
  prefix_length = 64

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # from_pool = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `from_pool` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_pool6`.
* `prefix` - (Required) Type: `string`. Default: `fd00:db8::/56`.
* `prefix_length` - (Required) Type: `int`. Default: `64`.

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
