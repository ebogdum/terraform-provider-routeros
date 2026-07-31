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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `chain` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `do` - (Read-only) Type: `string`.
* `do_group_num` - (Optional) Type: `string`. RouterOS `do-group-num`.
* `do_group_prfx` - (Optional) Type: `string`. RouterOS `do-group-prfx`.
* `do_jump` - (Optional) Type: `string`. RouterOS `do-jump`.
* `do_select_num` - (Optional) Type: `string`. RouterOS `do-select-num`.
* `do_select_prfx` - (Optional) Type: `string`. RouterOS `do-select-prfx`.
* `do_take` - (Optional) Type: `string`. RouterOS `do-take`.
* `do_where` - (Optional) Type: `string`. RouterOS `do-where`.
* `invalid` - (Read-only) Type: `bool`.
* `type` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
