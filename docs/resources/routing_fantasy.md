---
page_title: "RouterOS: routeros_routing_fantasy"
description: |-
  RouterOS resource.
---

# Resource: routeros_routing_fantasy

Manages the RouterOS `/routing/fantasy` menu.

## Example Usage

```terraform
resource "routeros_routing_fantasy" "fantasy_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # dst_address = "10.99.0.0/24"
  # gateway = "replace-me"
  # name = "tf-example"
  # scope = "replace-me"
  # target_scope = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dst_address` - (Optional) Type: `string`.
* `gateway` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `scope` - (Optional) Type: `string`.
* `target_scope` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_fantasy.example '*3'

# Named router
terraform import routeros_routing_fantasy.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_fantasy.example 'home/my-resource-name'
```
