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
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `chain` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.

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
