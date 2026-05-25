---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_rpki"
description: |-
  RouterOS resource.
---

# Resource: routeros_routing_rpki

Manages the RouterOS `/routing/rpki` menu.

## Example Usage

```terraform
resource "routeros_routing_rpki" "rpki_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "10.99.0.1"
  # expire_interval = "replace-me"
  # group = "replace-me"
  # port = "443"
  # preference = "replace-me"
  # refresh_interval = "replace-me"
  # retry_interval = "replace-me"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `expire_interval` - (Optional) Type: `string`.
* `group` - (Optional) Type: `string`.
* `port` - (Optional) Type: `string`.
* `preference` - (Optional) Type: `string`.
* `refresh_interval` - (Optional) Type: `string`.
* `retry_interval` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_rpki.example '*3'

# Named router
terraform import routeros_routing_rpki.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_rpki.example 'home/my-resource-name'
```
