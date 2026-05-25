---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_graphing_queue"
description: |-
  Same as graphing/interface — skipped to avoid "already exists" on re-runs.
---

# Resource: routeros_tool_graphing_queue

Same as graphing/interface — skipped to avoid "already exists" on re-runs.

## Example Usage

```terraform
resource "routeros_tool_graphing_queue" "queue_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_tool_graphing_queue.example '*3'

# Named router
terraform import routeros_tool_graphing_queue.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_tool_graphing_queue.example 'home/my-resource-name'
```
