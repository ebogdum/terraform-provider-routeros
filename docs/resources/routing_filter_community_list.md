---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_filter_community_list"
description: |-
  7.x routing-filter community-list uses 'communities' field and rule chain semantics that vary across releases. Skipped.
---

# Resource: routeros_routing_filter_community_list

7.x routing-filter community-list uses 'communities' field and rule chain semantics that vary across releases. Skipped.

## Example Usage

```terraform
resource "routeros_routing_filter_community_list" "community_list_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `communities` - (Optional) Type: `string`. RouterOS `communities`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `list` - (Optional) Type: `string`. RouterOS `list`.
* `regexp` - (Optional) Type: `string`. RouterOS `regexp`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_filter_community_list.example '*3'

# Named router
terraform import routeros_routing_filter_community_list.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_filter_community_list.example 'home/my-resource-name'
```
