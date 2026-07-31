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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `auto_upgrade` - (Optional) Type: `string`.
* `baud_rate` - (Optional) Type: `string`. RouterOS `baud-rate`.
* `boot_delay` - (Optional) Type: `string`. RouterOS `boot-delay`.
* `boot_device` - (Optional) Type: `string`.
* `boot_os` - (Optional) Type: `string`. RouterOS `boot-os`.
* `boot_protocol` - (Optional) Type: `string`.
* `cpu_frequency` - (Optional) Type: `string`.
* `disable_pci` - (Optional) Type: `string`. RouterOS `disable-pci`.
* `enable_jumper_reset` - (Optional) Type: `string`. RouterOS `enable-jumper-reset`.
* `enter_setup_on` - (Optional) Type: `string`. RouterOS `enter-setup-on`.
* `etherboot_port` - (Optional) Type: `string`. RouterOS `etherboot-port`.
* `force_backup_booter` - (Optional) Type: `string`.
* `gpio_function` - (Optional) Type: `string`. RouterOS `gpio-function`.
* `init_delay` - (Optional) Type: `string`. RouterOS `init-delay`.
* `preboot_etherboot` - (Optional) Type: `string`.
* `preboot_etherboot_server` - (Optional) Type: `string`.
* `preferred_architecture` - (Optional) Type: `string`. RouterOS `preferred-architecture`.
* `protected_routerboot` - (Optional) Type: `string`.
* `reformat_hold_button` - (Optional) Type: `string`.
* `reformat_hold_button_max` - (Optional) Type: `string`.
* `silent_boot` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_routerboard_settings.this 'home'
```
