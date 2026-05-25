---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_pool"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_pool

Manages the RouterOS `/ip/pool` menu.

## Example Usage

```terraform
data "routeros_ip_pool" "pool_example" {
  # router   = "my-router"  # omit for the default router
  # filter   = { name = "some-name" }
  # proplist = ["name", "address"]
}
```

## Argument Reference

This data source supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to query.
* `filter` - (Optional) Map of field=value pairs to narrow the result set.
* `proplist` - (Optional) List of property names to project; smaller payload.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `name` - (Required) Type: `string`. Unique identifier of the pool. Default: `tf_acc_pool`.
* `next_pool` - (Optional) Type: `string`. When IP address acquisition is performed a pool that has no free addresses, and the next-pool property is set, then IP address will be acquired from next-pool.
* `ranges` - (Required) Type: `string`. IP address list of non-overlapping IP address ranges in the form of: from1-to1,from2-to2,...,fromN-toN. For example, 10.0.0.1-10.0.0.27,10.0.0.32-10.0.0.47. Default: `10.255.255.0-10.255.255.4`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

