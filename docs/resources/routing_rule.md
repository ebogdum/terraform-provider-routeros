---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_rule"
description: |-
  RouterOS resource.
---

# Resource: routeros_routing_rule

Manages the RouterOS `/routing/rule` menu.

## Example Usage

```terraform
resource "routeros_routing_rule" "rule_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "lookup"
  # chain = "replace-me"
  # dst_address = "10.99.0.0/24"
  # interface = "ether1"
  # min_prefix = "replace-me"
  # realm = "replace-me"
  # routing_mark = "replace-me"
  # src_address = "10.99.0.0/24"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `action` - (Optional) Type: `enum(lookup|unreachable|drop|lookup-only-in-table|mangle)`.
* `chain` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dst_address` - (Optional) Type: `string`.
* `interface` - (Optional) Type: `string`.
* `min_prefix` - (Optional) Type: `string`.
* `realm` - (Optional) Type: `string`.
* `routing_mark` - (Optional) Type: `string`.
* `src_address` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `bool`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `invalid` - Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_rule.example '*3'

# Named router
terraform import routeros_routing_rule.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_rule.example 'home/my-resource-name'
```
