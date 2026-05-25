---
subcategory: "System & misc"
page_title: "RouterOS: routeros_caps_man_aaa"
description: |-
  RouterOS resource.
---

# Resource: routeros_caps_man_aaa

Manages the RouterOS `/caps-man/aaa` menu.

## Example Usage

```terraform
resource "routeros_caps_man_aaa" "aaa_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # called_format = "replace-me"
  # interim_update = "replace-me"
  # mac_caching = "replace-me"
  # mac_format = "replace-me"
  # mac_mode = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `called_format` - (Optional) Type: `string`.
* `interim_update` - (Optional) Type: `string`.
* `mac_caching` - (Optional) Type: `string`.
* `mac_format` - (Optional) Type: `string`.
* `mac_mode` - (Optional) Type: `string`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_caps_man_aaa.this 'home'
```
