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
  # acquired = false
  # block_device = false
  # btrfs = "replace-me"
  # eject_drive = "replace-me"
  # empty = false
  # encrypted = false
  # formatting = false
  # guid_partition_table = false
  # i_scsi_export = false
  # i_scsi_server_iqn = "replace-me"
  # i_scsi_server_port = "443"
  # iscsi_sharing = "replace-me"
  # media_interface = "replace-me"
  # media_sharing = false
  # mount_compress = false
  # mount_filesystem = true
  # mount_point_template = "replace-me"
  # mount_read_only = false
  # mounted = false
  # newfileman = "replace-me"
  # nfs_sharing = false
  # nvme = "replace-me"
  # nvme_tcp_export = false
  # nvme_tcp_server_allow_host_name = "replace-me"
  # nvme_tcp_server_nqn = "replace-me"
  # nvme_tcp_server_password = "REDACTED"
  # nvme_tcp_server_port = "443"
  # oldfileman = "replace-me"
  # parent = "4.294967295e+09"
  # part = "replace-me"
  # partition = false
  # partition_offset = "65536"
  # partition_size = "replace-me"
  # raid = "replace-me"
  # raid_and_master = "replace-me"
  # raid_and_type = "replace-me"
  # raid_master = "4.294967295e+09"
  # raid_member = false
  # raid_member_failed = false
  # raid_role = "spare"
  # raid_scrub = "replace-me"
  # read_only = false
  # reset_counters = "replace-me"
  # rose = "replace-me"
  # scan = "replace-me"
  # self_encrypted_and_locked = false
  # self_encryption_enabled = false
  # self_encryption_password = "REDACTED"
  # self_encryption_supported = false
  # slot = "replace-me"
  # smb_server_encryption = false
  # smb_server_password = "REDACTED"
  # smb_server_user = "replace-me"
  # smb_sharing = false
  # swap = false
  # swap_enabled = false
  # tmpfs = "replace-me"
  # tmpfs_max_size = "replace-me"
  # trim = "replace-me"
  # type = "6"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `acquired` - (Optional) Type: `bool`.
* `active_time` - (Read-only) Type: `string`.
* `available_spare` - (Read-only) Type: `int`.
* `available_spare_threshold` - (Read-only) Type: `int`.
* `block_device` - (Optional) Type: `bool`.
* `btrfs` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `controller_burst_time` - (Read-only) Type: `string`.
* `critical_temperature` - (Read-only) Type: `int`.
* `critical_temperature_time` - (Read-only) Type: `string`.
* `critical_warning` - (Read-only) Type: `string`.
* `default_slot` - (Read-only) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `discard_bytes` - (Read-only) Type: `string`.
* `discard_merges` - (Read-only) Type: `string`.
* `discard_ops` - (Read-only) Type: `string`.
* `discard_time` - (Read-only) Type: `string`.
* `eject_drive` - (Optional) Type: `string`.
* `empty` - (Optional) Type: `bool`.
* `encrypted` - (Optional) Type: `bool`.
* `flush_ops` - (Read-only) Type: `string`.
* `flush_time` - (Read-only) Type: `string`.
* `formatting` - (Optional) Type: `bool`.
* `free` - (Read-only) Type: `string`.
* `fs` - (Read-only) Type: `string`.
* `fw_version` - (Read-only) Type: `string`.
* `guid_partition_table` - (Optional) Type: `bool`.
* `host_read_bytes` - (Read-only) Type: `string`.
* `host_read_commands` - (Read-only) Type: `string`.
* `host_write_bytes` - (Read-only) Type: `string`.
* `host_write_commands` - (Read-only) Type: `string`.
* `i_scsi_export` - (Optional) Type: `bool`.
* `i_scsi_server_iqn` - (Optional) Type: `string`.
* `i_scsi_server_port` - (Optional) Type: `int`.
* `in_flight_ops` - (Read-only) Type: `string`.
* `interface` - (Read-only) Type: `string`.
* `interface_speed` - (Read-only) Type: `string`.
* `iscsi_sharing` - (Optional) Type: `string`.
* `label` - (Read-only) Type: `string`.
* `media_interface` - (Optional) Type: `string`.
* `media_sharing` - (Optional) Type: `bool`.
* `model` - (Read-only) Type: `string`.
* `mount_compress` - (Optional) Type: `bool`.
* `mount_filesystem` - (Optional) Type: `bool`.
* `mount_point` - (Read-only) Type: `string`.
* `mount_point_template` - (Optional) Type: `string`. Sets the mounting point for the file system. It is possible to set the mount point as the following parameters based on the disk: [slot] (default) - sets the mount point as the slot name. [model] - sets the mount point as the device's model name. [serial] - sets the mount point as the device serial [fw-version] - sets the mount point as the device's firmware version. [fs-label] - sets the mount point as the device's file system label. [fs-uuid] - sets the mount point as the device's UUID [fs] - sets the mount point as the device's file system ros Additionally, it is possible to combine multiple variables to create a single mount point: ros
* `mount_read_only` - (Optional) Type: `bool`. Sets the mounted disk in read only mode when set to yes .
* `mounted` - (Optional) Type: `bool`.
* `newfileman` - (Optional) Type: `string`.
* `nfs_sharing` - (Optional) Type: `bool`.
* `nvme` - (Optional) Type: `string`.
* `nvme_tcp_export` - (Optional) Type: `bool`.
* `nvme_tcp_server_allow_host_name` - (Optional) Type: `string`.
* `nvme_tcp_server_nqn` - (Optional) Type: `string`.
* `nvme_tcp_server_password` - (Optional) Type: `string`. **Sensitive.**
* `nvme_tcp_server_port` - (Optional) Type: `int`.
* `nvme_tcp_server_secret` - (Read-only) Type: `string`. **Sensitive.**
* `oldfileman` - (Optional) Type: `string`.
* `parent` - (Optional) Type: `string`.
* `part` - (Optional) Type: `string`.
* `partition` - (Optional) Type: `bool`.
* `partition_number` - (Read-only) Type: `int`.
* `partition_offset` - (Optional) Type: `string`.
* `partition_size` - (Optional) Type: `string`.
* `percentage_used` - (Read-only) Type: `int`.
* `power_cycles` - (Read-only) Type: `int`.
* `power_on_time` - (Read-only) Type: `string`.
* `raid` - (Optional) Type: `string`.
* `raid_and_master` - (Optional) Type: `string`.
* `raid_and_type` - (Optional) Type: `string`.
* `raid_master` - (Optional) Type: `string`.
* `raid_member` - (Optional) Type: `bool`.
* `raid_member_failed` - (Optional) Type: `bool`.
* `raid_role` - (Optional) Type: `string`.
* `raid_scrub` - (Optional) Type: `string`.
* `read_bytes` - (Read-only) Type: `string`.
* `read_merges` - (Read-only) Type: `string`.
* `read_only` - (Optional) Type: `bool`.
* `read_ops` - (Read-only) Type: `string`.
* `read_ops_per_second` - (Read-only) Type: `string`.
* `read_rate` - (Read-only) Type: `string`.
* `read_time` - (Read-only) Type: `string`.
* `reset_counters` - (Optional) Type: `string`.
* `rose` - (Optional) Type: `string`.
* `scan` - (Optional) Type: `string`.
* `self_encrypted_and_locked` - (Optional) Type: `bool`.
* `self_encryption_enabled` - (Optional) Type: `bool`.
* `self_encryption_password` - (Optional) Type: `string`. **Sensitive.**
* `self_encryption_supported` - (Optional) Type: `bool`.
* `serial` - (Read-only) Type: `string`.
* `size` - (Read-only) Type: `string`.
* `slot` - (Optional) Type: `string`.
* `smb_server_encryption` - (Optional) Type: `bool`.
* `smb_server_password` - (Optional) Type: `string`. **Sensitive.**
* `smb_server_user` - (Optional) Type: `string`.
* `smb_sharing` - (Optional) Type: `bool`.
* `state` - (Read-only) Type: `string`.
* `swap` - (Optional) Type: `bool`.
* `swap_enabled` - (Optional) Type: `bool`.
* `temperature` - (Read-only) Type: `int`.
* `temperatures` - (Read-only) Type: `string`.
* `tmpfs` - (Optional) Type: `string`.
* `tmpfs_max_size` - (Optional) Type: `string`.
* `trim` - (Optional) Type: `string`.
* `type` - (Optional) Type: `string`.
* `unrecovered_integrity_errors` - (Read-only) Type: `int`.
* `unsafe_shutdown` - (Read-only) Type: `int`.
* `use` - (Read-only) Type: `int`.
* `uuid` - (Read-only) Type: `string`.
* `wait_time` - (Read-only) Type: `string`.
* `warning_temperature` - (Read-only) Type: `int`.
* `warning_temperature_time` - (Read-only) Type: `string`.
* `write_bytes` - (Read-only) Type: `string`.
* `write_merges` - (Read-only) Type: `string`.
* `write_ops` - (Read-only) Type: `string`.
* `write_ops_per_second` - (Read-only) Type: `string`.
* `write_rate` - (Read-only) Type: `string`.
* `write_time` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
