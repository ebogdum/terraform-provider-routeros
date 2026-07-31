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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `called_format` - (Optional) Type: `string`.
* `interim_update` - (Optional) Type: `string`.
* `mac_caching` - (Optional) Type: `string`.
* `mac_format` - (Optional) Type: `string`.
* `mac_mode` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_caps_man_aaa.this 'home'
```
