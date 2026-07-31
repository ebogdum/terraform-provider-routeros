---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_table"
description: |-
  RouterOS resource.
---

# Resource: routeros_routing_table

Manages the RouterOS `/routing/table` menu.

## Example Usage

```terraform
resource "routeros_routing_table" "table_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # fib = false
  # name = "tf-example"
  # used = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dynamic` - (Read-only) Type: `bool`.
* `fib` - (Optional) Type: `bool`.
* `invalid` - (Read-only) Type: `bool`.
* `name` - (Optional) Type: `string`.
* `usage` - (Read-only) Type: `int`.
* `used` - (Read-only) Type: `bool`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_table.example '*3'

# Named router
terraform import routeros_routing_table.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_table.example 'home/my-resource-name'
```
