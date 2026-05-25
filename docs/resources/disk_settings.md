---
page_title: "RouterOS: routeros_disk_settings"
description: |-
  RouterOS resource.
---

# Resource: routeros_disk_settings

Manages the RouterOS `/disk/settings` menu.

## Example Usage

```terraform
resource "routeros_disk_settings" "settings_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # auto_media_interface = "replace-me"
  # auto_media_sharing = false
  # auto_smb_sharing = false
  # auto_smb_user = "replace-me"
  # default_mount_point_template = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `auto_media_interface` - (Optional) Type: `string`.
* `auto_media_sharing` - (Optional) Type: `bool`.
* `auto_smb_sharing` - (Optional) Type: `bool`.
* `auto_smb_user` - (Optional) Type: `string`.
* `default_mount_point_template` - (Optional) Type: `string`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_disk_settings.this 'home'
```
