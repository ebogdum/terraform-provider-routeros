---
page_title: "RouterOS: routeros_ip_pool"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_pool

Manages the RouterOS `/ip/pool` menu.

## Example Usage

```terraform
resource "routeros_ip_pool" "pool_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"
  ranges = "10.99.0.100-10.99.0.200"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # next_pool = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `name` - (Required) Type: `string`. Unique identifier of the pool. Default: `tf_acc_pool`.
* `next_pool` - (Optional) Type: `string`. When IP address acquisition is performed a pool that has no free addresses, and the next-pool property is set, then IP address will be acquired from next-pool.
* `ranges` - (Required) Type: `string`. IP address list of non-overlapping IP address ranges in the form of: from1-to1,from2-to2,...,fromN-toN. For example, 10.0.0.1-10.0.0.27,10.0.0.32-10.0.0.47. Default: `10.255.255.0-10.255.255.4`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_pool.example '*3'

# Named router
terraform import routeros_ip_pool.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_pool.example 'home/my-resource-name'
```
