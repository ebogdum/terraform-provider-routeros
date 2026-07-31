---
subcategory: "Certificates"
page_title: "RouterOS: routeros_certificate_settings"
description: |-
  RouterOS resource.
---

# Resource: routeros_certificate_settings

Manages the RouterOS `/certificate/settings` menu.

## Example Usage

```terraform
resource "routeros_certificate_settings" "settings_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # builtin_trust_store = "replace-me"
  # crl_download = false
  # crl_store = "replace-me"
  # crl_use = false
  # current_defaults = []
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `builtin_trust_store` - (Optional) Type: `string`.
* `crl_download` - (Optional) Type: `bool`.
* `crl_store` - (Optional) Type: `string`.
* `crl_use` - (Optional) Type: `bool`.
* `current_defaults` - (Optional) Type: `list`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_certificate_settings.this 'home'
```
