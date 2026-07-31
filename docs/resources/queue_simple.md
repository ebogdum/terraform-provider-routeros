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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `bucket_size` - (Optional) Type: `string`. RouterOS `bucket-size`.
* `burst_limit` - (Optional) Type: `string`.
* `burst_threshold` - (Optional) Type: `string`.
* `burst_time` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dst` - (Optional) Type: `string`. RouterOS `dst`.
* `limit_at` - (Optional) Type: `string`.
* `max_limit` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `packet_marks` - (Optional) Type: `string`.
* `parent` - (Optional) Type: `string`.
* `place_before` - (Read-only) Type: `string`. RouterOS .id (e.g. *3) of the entry this one should be moved before. Use to enforce explicit ordering on ordered menus.
* `place_before_ros` - (Optional) Type: `string`.
* `priority` - (Optional) Type: `string`.
* `queue` - (Optional) Type: `string`.
* `target` - (Required) Type: `string`.
* `time` - (Optional) Type: `string`. RouterOS `time`.
* `total_bucket_size` - (Optional) Type: `string`. RouterOS `total-bucket-size`.
* `total_burst_limit` - (Optional) Type: `string`. RouterOS `total-burst-limit`.
* `total_burst_threshold` - (Optional) Type: `string`. RouterOS `total-burst-threshold`.
* `total_burst_time` - (Optional) Type: `string`. RouterOS `total-burst-time`.
* `total_limit_at` - (Optional) Type: `string`. RouterOS `total-limit-at`.
* `total_max_limit` - (Optional) Type: `string`. RouterOS `total-max-limit`.
* `total_priority` - (Optional) Type: `string`. RouterOS `total-priority`.
* `total_queue` - (Optional) Type: `string`. RouterOS `total-queue`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
