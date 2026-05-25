---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_bgp_template"
description: |-
  RouterOS resource.
---

# Resource: routeros_routing_bgp_template

Manages the RouterOS `/routing/bgp/template` menu.

## Example Usage

```terraform
resource "routeros_routing_bgp_template" "template_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # afi = "replace-me"
  # as = "65000"
  # hold_time = "replace-me"
  # keepalive_time = "replace-me"
  # multihop = "replace-me"
  # nexthop_choice = "replace-me"
  # router_id = "1.1.1.1"
  # routing_table = "main"
  # templates = "replace-me"
  # use_bfd = "replace-me"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `afi` - (Optional) Type: `string`.
* `as` - (Optional) Type: `string`. Default: `65000`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `hold_time` - (Optional) Type: `string`.
* `keepalive_time` - (Optional) Type: `string`.
* `multihop` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_bgptpl`.
* `nexthop_choice` - (Optional) Type: `string`.
* `router_id` - (Optional) Type: `string`. Default: `1.1.1.1`.
* `routing_table` - (Optional) Type: `string`.
* `templates` - (Optional) Type: `string`.
* `use_bfd` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `default` - Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_bgp_template.example '*3'

# Named router
terraform import routeros_routing_bgp_template.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_bgp_template.example 'home/my-resource-name'
```
