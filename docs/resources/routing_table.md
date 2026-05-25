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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `fib` - (Optional) Type: `bool`.
* `name` - (Optional) Type: `string`.
* `used` - (Optional) Type: `bool`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.
* `usage` - Type: `int`.

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
