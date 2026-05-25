---
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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `builtin_trust_store` - (Optional) Type: `string`.
* `crl_download` - (Optional) Type: `bool`.
* `crl_store` - (Optional) Type: `string`.
* `crl_use` - (Optional) Type: `bool`.
* `current_defaults` - (Optional) Type: `list`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_certificate_settings.this 'home'
```
