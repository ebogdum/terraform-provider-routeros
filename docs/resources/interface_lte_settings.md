---
subcategory: "Interface"
page_title: "RouterOS: routeros_interface_lte_settings"
description: |-
  Mirrors RouterOS /interface/lte/settings.
---

# Resource: routeros_interface_lte_settings

Mirrors RouterOS `/interface/lte/settings`.

## Example Usage

```terraform
resource "routeros_interface_lte_settings" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # esim_channel = "replace-me"
  # firmware_path = "replace-me"
  # link_recovery_timer = 0
  # mode = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `esim_channel` - (Optional) Type: `string`. RouterOS `esim-channel`.
* `firmware_path` - (Optional) Type: `string`. RouterOS `firmware-path`.
* `link_recovery_timer` - (Optional) Type: `int`. RouterOS `link-recovery-timer`.
* `mode` - (Optional) Type: `string`. RouterOS `mode`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_interface_lte_settings.this 'home'
```
