---
subcategory: "MPLS"
page_title: "RouterOS: routeros_mpls_settings"
description: |-
  RouterOS resource.
---

# Resource: routeros_mpls_settings

Manages the RouterOS `/mpls/settings` menu.

## Example Usage

```terraform
resource "routeros_mpls_settings" "settings_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # allow_fast_path = false
  # dynamic_label_range = "replace-me"
  # propagate_ttl = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `allow_fast_path` - (Optional) Type: `bool`.
* `dynamic_label_range` - (Optional) Type: `string`.
* `mpls_fast_path_bytes` - (Optional) Type: `int`.
* `mpls_fast_path_packets` - (Optional) Type: `int`.
* `propagate_ttl` - (Optional) Type: `bool`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_mpls_settings.this 'home'
```
