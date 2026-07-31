---
subcategory: "Routing"
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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dealer_id` - (Optional) Type: `string`. RouterOS `dealer-id`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dst_address` - (Optional) Type: `string`.
* `gateway` - (Optional) Type: `string`.
* `instance_id` - (Optional) Type: `string`. RouterOS `instance-id`.
* `name` - (Optional) Type: `string`.
* `offset` - (Optional) Type: `string`. RouterOS `offset`.
* `prefix_length` - (Optional) Type: `string`. RouterOS `prefix-length`.
* `priv_offs` - (Optional) Type: `string`. RouterOS `priv-offs`.
* `priv_size` - (Optional) Type: `string`. RouterOS `priv-size`.
* `route_count` - (Optional) Type: `string`. RouterOS `count`.
* `scope` - (Optional) Type: `string`.
* `seed` - (Optional) Type: `string`. RouterOS `seed`.
* `target_scope` - (Optional) Type: `string`.
* `use_hold` - (Optional) Type: `string`. RouterOS `use-hold`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
