---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_filter_select_rule"
description: |-
  References a /routing/filter rule. Skipped.
---

# Resource: routeros_routing_filter_select_rule

References a /routing/filter rule. Skipped.

## Example Usage

```terraform
resource "routeros_routing_filter_select_rule" "select_rule_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # chain = "replace-me"
  # type = "where"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `chain` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `type` - (Optional) Type: `enum(where|group-num|group-prfx|select-num|select-prfx|take, ...)`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `do` - Type: `string`.
* `invalid` - Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_filter_select_rule.example '*3'

# Named router
terraform import routeros_routing_filter_select_rule.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_filter_select_rule.example 'home/my-resource-name'
```
