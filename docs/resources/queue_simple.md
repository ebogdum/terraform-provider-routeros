---
subcategory: "Queues"
page_title: "RouterOS: routeros_queue_simple"
description: |-
  RouterOS resource.
---

# Resource: routeros_queue_simple

Manages the RouterOS `/queue/simple` menu.

## Example Usage

```terraform
resource "routeros_queue_simple" "simple_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"
  target = "127.0.0.1/32"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # burst_limit = "replace-me"
  # burst_threshold = "replace-me"
  # burst_time = "replace-me"
  # limit_at = "replace-me"
  # max_limit = "replace-me"
  # packet_marks = "replace-me"
  # parent = "replace-me"
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
* `name` - (Required) Type: `string`. Default: `tf_acc_queue`.
* `packet_marks` - (Optional) Type: `string`.
* `parent` - (Optional) Type: `string`.
* `place_before` - (Optional) Type: `string`.
* `priority` - (Optional) Type: `string`.
* `queue` - (Optional) Type: `string`.
* `target` - (Required) Type: `string`. Default: `127.0.0.1/32`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_queue_simple.example '*3'

# Named router
terraform import routeros_queue_simple.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_queue_simple.example 'home/my-resource-name'
```
