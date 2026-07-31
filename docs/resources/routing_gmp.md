---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_gmp"
description: |-
  GMP needs interface configuration. Skipped from automated acc tests.
---

# Resource: routeros_routing_gmp

GMP needs interface configuration. Skipped from automated acc tests.

## Example Usage

```terraform
resource "routeros_routing_gmp" "gmp_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # exclude = false
  # group = "replace-me"
  # interfaces = "replace-me"
  # sources = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `exclude` - (Optional) Type: `bool`.
* `group` - (Read-only) Type: `string`.
* `groups` - (Optional) Type: `string`. RouterOS `groups`.
* `interfaces` - (Optional) Type: `string`.
* `sources` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_gmp.example '*3'

# Named router
terraform import routeros_routing_gmp.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_gmp.example 'home/my-resource-name'
```
