---
subcategory: "Queues"
page_title: "RouterOS: routeros_queue_tree"
description: |-
  RouterOS resource.
---

# Resource: routeros_queue_tree

Manages the RouterOS `/queue/tree` menu.

## Example Usage

```terraform
resource "routeros_queue_tree" "tree_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"
  parent = "global"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # burst_limit = "replace-me"
  # burst_threshold = "replace-me"
  # burst_time = "replace-me"
  # limit_at = "replace-me"
  # max_limit = "replace-me"
  # packet_mark = "replace-me"
  # place_before = "replace-me"
  # priority = "replace-me"
  # queue = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
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

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_queue_tree.example '*3'

# Named router
terraform import routeros_queue_tree.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_queue_tree.example 'home/my-resource-name'
```
