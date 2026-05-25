---
page_title: "RouterOS: routeros_queue_tree"
description: |-
  RouterOS resource.
---

# Data Source: routeros_queue_tree

Manages the RouterOS `/queue/tree` menu.

## Example Usage

```terraform
data "routeros_queue_tree" "tree_example" {
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
* `burst_limit` - (Optional) Type: `string`.
* `burst_threshold` - (Optional) Type: `string`.
* `burst_time` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `limit_at` - (Optional) Type: `string`.
* `max_limit` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_qtree`.
* `packet_mark` - (Optional) Type: `string`.
* `parent` - (Required) Type: `string`. Parent queue ("global" for top-level, or another queue's name/id). Default: `global`.
* `place_before` - (Optional) Type: `string`.
* `priority` - (Optional) Type: `string`.
* `queue` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

