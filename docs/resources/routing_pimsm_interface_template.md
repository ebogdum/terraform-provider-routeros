---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_pimsm_interface_template"
description: |-
  Discovered; needs pimsm instance
---

# Resource: routeros_routing_pimsm_interface_template

Discovered; needs pimsm instance

## Example Usage

```terraform
resource "routeros_routing_pimsm_interface_template" "interface_template_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # hello_delay = "replace-me"
  # hello_period = "replace-me"
  # instance = "replace-me"
  # interfaces = "replace-me"
  # join_prune_period = "replace-me"
  # join_tracking_support = "replace-me"
  # override_interval = "replace-me"
  # priority = "replace-me"
  # propagation_delay = "replace-me"
  # source_addresses = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `disabled` - (Optional) Type: `bool`.
* `hello_delay` - (Optional) Type: `string`.
* `hello_period` - (Optional) Type: `string`.
* `instance` - (Optional) Type: `string`.
* `interfaces` - (Optional) Type: `string`.
* `join_prune_period` - (Optional) Type: `string`.
* `join_tracking_support` - (Optional) Type: `string`.
* `override_interval` - (Optional) Type: `string`.
* `priority` - (Optional) Type: `string`.
* `propagation_delay` - (Optional) Type: `string`.
* `source_addresses` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_pimsm_interface_template.example '*3'

# Named router
terraform import routeros_routing_pimsm_interface_template.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_pimsm_interface_template.example 'home/my-resource-name'
```
