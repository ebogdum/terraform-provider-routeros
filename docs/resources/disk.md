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
  # media_interface = "replace-me"
  # media_sharing = false
  # mount_compress = false
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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `acquired` - (Optional) Type: `bool`.
* `active_time` - (Read-only) Type: `string`.
* `available_spare` - (Read-only) Type: `int`.
* `available_spare_threshold` - (Read-only) Type: `int`.
* `block_device` - (Read-only) Type: `bool`.
* `btrfs` - (Read-only) Type: `string`.
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
* `eject_drive` - (Read-only) Type: `string`.
* `empty` - (Read-only) Type: `bool`.
* `encrypted` - (Read-only) Type: `bool`.
* `flush_ops` - (Read-only) Type: `string`.
* `flush_time` - (Read-only) Type: `string`.
* `formatting` - (Read-only) Type: `bool`.
* `free` - (Read-only) Type: `string`.
* `fs` - (Read-only) Type: `string`.
* `fw_version` - (Read-only) Type: `string`.
* `guid_partition_table` - (Read-only) Type: `bool`.
* `host_read_bytes` - (Read-only) Type: `string`.
* `host_read_commands` - (Read-only) Type: `string`.
* `host_write_bytes` - (Read-only) Type: `string`.
* `host_write_commands` - (Read-only) Type: `string`.
* `i_scsi_export` - (Read-only) Type: `bool`.
* `i_scsi_server_iqn` - (Read-only) Type: `string`.
* `i_scsi_server_port` - (Read-only) Type: `int`.
* `in_flight_ops` - (Read-only) Type: `string`.
* `interface` - (Read-only) Type: `string`.
* `interface_speed` - (Read-only) Type: `string`.
* `iscsi_sharing` - (Read-only) Type: `string`.
* `label` - (Read-only) Type: `string`.
* `media_interface` - (Optional) Type: `string`.
* `media_sharing` - (Optional) Type: `bool`.
* `model` - (Read-only) Type: `string`.
* `mount_compress` - (Optional) Type: `bool`. Filesystem compression. RouterOS calls this `compress`.
* `mount_filesystem` - (Optional) Type: `bool`.
* `mount_point` - (Read-only) Type: `string`.
* `mount_point_template` - (Optional) Type: `string`. Sets the mounting point for the file system. It is possible to set the mount point as the following parameters based on the disk: [slot] (default) - sets the mount point as the slot name. [model] - sets the mount point as the device's model name. [serial] - sets the mount point as the device serial [fw-version] - sets the mount point as the device's firmware version. [fs-label] - sets the mount point as the device's file system label. [fs-uuid] - sets the mount point as the device's UUID [fs] - sets the mount point as the device's file system ros Additionally, it is possible to combine multiple variables to create a single mount point: ros
* `mount_read_only` - (Optional) Type: `bool`. Sets the mounted disk in read only mode when set to yes .
* `mounted` - (Read-only) Type: `bool`.
* `newfileman` - (Read-only) Type: `string`.
* `nfs_sharing` - (Read-only) Type: `bool`.
* `nvme` - (Read-only) Type: `string`.
* `nvme_tcp_export` - (Read-only) Type: `bool`.
* `nvme_tcp_server_allow_host_name` - (Read-only) Type: `string`.
* `nvme_tcp_server_nqn` - (Read-only) Type: `string`.
* `nvme_tcp_server_password` - (Read-only) Type: `string`. **Sensitive.**
* `nvme_tcp_server_port` - (Read-only) Type: `int`.
* `nvme_tcp_server_secret` - (Read-only) Type: `string`. **Sensitive.**
* `oldfileman` - (Read-only) Type: `string`.
* `parent` - (Optional) Type: `string`.
* `part` - (Read-only) Type: `string`.
* `partition` - (Read-only) Type: `bool`.
* `partition_number` - (Read-only) Type: `int`.
* `partition_offset` - (Optional) Type: `string`.
* `partition_size` - (Optional) Type: `string`.
* `percentage_used` - (Read-only) Type: `int`.
* `power_cycles` - (Read-only) Type: `int`.
* `power_on_time` - (Read-only) Type: `string`.
* `raid` - (Read-only) Type: `string`.
* `raid_and_master` - (Read-only) Type: `string`.
* `raid_and_type` - (Read-only) Type: `string`.
* `raid_master` - (Read-only) Type: `string`.
* `raid_member` - (Read-only) Type: `bool`.
* `raid_member_failed` - (Read-only) Type: `bool`.
* `raid_role` - (Read-only) Type: `string`.
* `raid_scrub` - (Read-only) Type: `string`.
* `read_bytes` - (Read-only) Type: `string`.
* `read_merges` - (Read-only) Type: `string`.
* `read_only` - (Read-only) Type: `bool`.
* `read_ops` - (Read-only) Type: `string`.
* `read_ops_per_second` - (Read-only) Type: `string`.
* `read_rate` - (Read-only) Type: `string`.
* `read_time` - (Read-only) Type: `string`.
* `reset_counters` - (Read-only) Type: `string`.
* `rose` - (Read-only) Type: `string`.
* `scan` - (Read-only) Type: `string`.
* `self_encrypted_and_locked` - (Read-only) Type: `bool`.
* `self_encryption_enabled` - (Read-only) Type: `bool`.
* `self_encryption_password` - (Read-only) Type: `string`. **Sensitive.**
* `self_encryption_supported` - (Read-only) Type: `bool`.
* `serial` - (Read-only) Type: `string`.
* `size` - (Read-only) Type: `string`.
* `slot` - (Optional) Type: `string`.
* `smb_server_encryption` - (Optional) Type: `bool`.
* `smb_server_password` - (Optional) Type: `string`. **Sensitive.**
* `smb_server_user` - (Optional) Type: `string`.
* `smb_sharing` - (Optional) Type: `bool`.
* `state` - (Read-only) Type: `string`.
* `swap` - (Optional) Type: `bool`.
* `swap_enabled` - (Read-only) Type: `bool`.
* `temperature` - (Read-only) Type: `int`.
* `temperatures` - (Read-only) Type: `string`.
* `tmpfs` - (Read-only) Type: `string`.
* `tmpfs_max_size` - (Optional) Type: `string`.
* `trim` - (Read-only) Type: `string`.
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
