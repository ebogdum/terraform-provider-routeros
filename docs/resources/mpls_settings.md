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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `allow_fast_path` - (Optional) Type: `bool`.
* `dynamic_label_range` - (Optional) Type: `string`.
* `propagate_ttl` - (Optional) Type: `bool`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `mpls_fast_path_bytes` - Type: `int`.
* `mpls_fast_path_packets` - Type: `int`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_mpls_settings.this 'home'
```
