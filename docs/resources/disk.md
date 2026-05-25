---
subcategory: "Storage"
page_title: "RouterOS: routeros_disk"
description: |-
  Storage volumes. Creating one usually requires a backing device.
---

# Resource: routeros_disk

Storage volumes. Creating one usually requires a backing device.

## Example Usage

```terraform
resource "routeros_disk" "disk_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # media_interface = "replace-me"
  # media_sharing = false
  # mount_filesystem = true
  # mount_point_template = "replace-me"
  # mount_read_only = false
  # parent = "4.294967295e+09"
  # partition_offset = "65536"
  # partition_size = "replace-me"
  # slot = "replace-me"
  # smb_server_encryption = false
  # smb_server_password = "REDACTED"
  # smb_server_user = "replace-me"
  # smb_sharing = false
  # swap = false
  # tmpfs_max_size = "replace-me"
  # type = "6"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `media_interface` - (Optional) Type: `string`.
* `media_sharing` - (Optional) Type: `bool`.
* `mount_filesystem` - (Optional) Type: `bool`. Default: `1`.
* `mount_point_template` - (Optional) Type: `string`. Sets the mounting point for the file system. It is possible to set the mount point as the following parameters based on the disk: [slot] (default) - sets the mount point as the slot name. [model] - sets the mount point as the device's model name. [serial] - sets the mount point as the device serial [fw-version] - sets the mount point as the device's firmware version. [fs-label] - sets the mount point as the device's file system label. [fs-uuid] - sets the mount point as the device's UUID [fs] - sets the mount point as the device's file system ros Additionally, it is possible to combine multiple variables to create a single mount point: ros.
* `mount_read_only` - (Optional) Type: `bool`. Sets the mounted disk in read only mode when set to yes .
* `parent` - (Optional) Type: `string`. Default: `4.294967295e+09`.
* `partition_offset` - (Optional) Type: `string`. Default: `65536`.
* `partition_size` - (Optional) Type: `string`.
* `slot` - (Optional) Type: `string`.
* `smb_server_encryption` - (Optional) Type: `bool`.
* `smb_server_password` - (Optional) Type: `string`. **Sensitive.**
* `smb_server_user` - (Optional) Type: `string`.
* `smb_sharing` - (Optional) Type: `bool`.
* `swap` - (Optional) Type: `bool`.
* `tmpfs_max_size` - (Optional) Type: `string`.
* `type` - (Optional) Type: `string`. Default: `6`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `partition_number` - Type: `int`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_disk.example '*3'

# Named router
terraform import routeros_disk.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_disk.example 'home/my-resource-name'
```
