---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_id"
description: |-
  RouterOS resource.
---

# Resource: routeros_routing_id

Manages the RouterOS `/routing/id` menu.

## Example Usage

```terraform
resource "routeros_routing_id" "id_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # name = "tf-example"
  # select_dynamic_id = "only-static"
  # select_from_vrf = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dynamic` - (Read-only) Type: `bool`.
* `dynamic_id` - (Read-only) Type: `string`.
* `inactive` - (Read-only) Type: `bool`.
* `invalid` - (Read-only) Type: `bool`.
* `name` - (Optional) Type: `string`.
* `select_dynamic_id` - (Optional) Type: `string`.
* `select_from_vrf` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_id.example '*3'

# Named router
terraform import routeros_routing_id.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_id.example 'home/my-resource-name'
```
