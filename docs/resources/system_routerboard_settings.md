---
subcategory: "System"
page_title: "RouterOS: routeros_system_routerboard_settings"
description: |-
  RouterOS resource.
---

# Resource: routeros_system_routerboard_settings

Manages the RouterOS `/system/routerboard/settings` menu.

## Example Usage

```terraform
resource "routeros_system_routerboard_settings" "settings_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # auto_upgrade = "replace-me"
  # boot_device = "replace-me"
  # boot_protocol = "replace-me"
  # cpu_frequency = "replace-me"
  # force_backup_booter = "replace-me"
  # preboot_etherboot = "replace-me"
  # preboot_etherboot_server = "replace-me"
  # protected_routerboot = "replace-me"
  # reformat_hold_button = "replace-me"
  # reformat_hold_button_max = "replace-me"
  # silent_boot = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `auto_upgrade` - (Optional) Type: `string`.
* `boot_device` - (Optional) Type: `string`.
* `boot_protocol` - (Optional) Type: `string`.
* `cpu_frequency` - (Optional) Type: `string`.
* `force_backup_booter` - (Optional) Type: `string`.
* `preboot_etherboot` - (Optional) Type: `string`.
* `preboot_etherboot_server` - (Optional) Type: `string`.
* `protected_routerboot` - (Optional) Type: `string`.
* `reformat_hold_button` - (Optional) Type: `string`.
* `reformat_hold_button_max` - (Optional) Type: `string`.
* `silent_boot` - (Optional) Type: `string`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_routerboard_settings.this 'home'
```
