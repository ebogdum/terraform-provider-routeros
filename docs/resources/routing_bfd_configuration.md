---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_bfd_configuration"
description: |-
  RouterOS resource.
---

# Resource: routeros_routing_bfd_configuration

Manages the RouterOS `/routing/bfd/configuration` menu.

## Example Usage

```terraform
resource "routeros_routing_bfd_configuration" "configuration_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address_list = "replace-me"
  # addresses = "replace-me"
  # forbid_bfd = "replace-me"
  # interfaces = "replace-me"
  # min_rx = "replace-me"
  # min_tx = "replace-me"
  # multiplier = "replace-me"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address_list` - (Optional) Type: `string`.
* `addresses` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `forbid_bfd` - (Optional) Type: `string`.
* `interfaces` - (Optional) Type: `string`.
* `min_rx` - (Optional) Type: `string`.
* `min_tx` - (Optional) Type: `string`.
* `multiplier` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_bfd_configuration.example '*3'

# Named router
terraform import routeros_routing_bfd_configuration.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_bfd_configuration.example 'home/my-resource-name'
```
