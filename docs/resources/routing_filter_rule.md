---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_filter_rule"
description: |-
  RouterOS resource.
---

# Resource: routeros_routing_filter_rule

Manages the RouterOS `/routing/filter/rule` menu.

## Example Usage

```terraform
resource "routeros_routing_filter_rule" "rule_example" {
  # router = "my-router"  # which router to target; omit for the default
  chain = "tf_acc_chain"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # rule = "accept"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `chain` - (Required) Type: `string`. Default: `tf_acc_chain`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `rule` - (Optional) Type: `string`. Default: `accept`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_filter_rule.example '*3'

# Named router
terraform import routeros_routing_filter_rule.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_filter_rule.example 'home/my-resource-name'
```
