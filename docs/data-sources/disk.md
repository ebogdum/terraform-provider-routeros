---
subcategory: "Storage"
page_title: "RouterOS: routeros_disk"
description: |-
  Storage volumes. Creating one usually requires a backing device.
---

# Data Source: routeros_disk

Storage volumes. Creating one usually requires a backing device.

## Example Usage

```terraform
data "routeros_disk" "disk_example" {
  # router   = "my-router"  # omit for the default router
  # filter   = { name = "some-name" }
  # proplist = ["name", "address"]
}
```

## Argument Reference

This data source supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to query.
* `filter` - (Optional) Map of field=value pairs to narrow the result set.
* `proplist` - (Optional) List of property names to project; smaller payload.
* `acquired` - (Optional) Type: `bool`.
* `block_device` - (Optional) Type: `bool`.
* `btrfs` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `eject_drive` - (Optional) Type: `string`.
* `empty` - (Optional) Type: `bool`.
* `encrypted` - (Optional) Type: `bool`.
* `formatting` - (Optional) Type: `bool`.
* `guid_partition_table` - (Optional) Type: `bool`.
* `i_scsi_export` - (Optional) Type: `bool`.
* `i_scsi_server_iqn` - (Optional) Type: `string`.
* `i_scsi_server_port` - (Optional) Type: `int`.
* `iscsi_sharing` - (Optional) Type: `string`.
* `media_interface` - (Optional) Type: `string`.
* `media_sharing` - (Optional) Type: `bool`.
* `mount_compress` - (Optional) Type: `bool`.
* `mount_filesystem` - (Optional) Type: `bool`. Default: `1`.
* `mount_point_template` - (Optional) Type: `string`. Sets the mounting point for the file system. It is possible to set the mount point as the following parameters based on the disk: [slot] (default) - sets the mount point as the slot name. [model] - sets the mount point as the device's model name. [serial] - sets the mount point as the device serial [fw-version] - sets the mount point as the device's firmware version. [fs-label] - sets the mount point as the device's file system label. [fs-uuid] - sets the mount point as the device's UUID [fs] - sets the mount point as the device's file system ros Additionally, it is possible to combine multiple variables to create a single mount point: ros.
* `mount_read_only` - (Optional) Type: `bool`. Sets the mounted disk in read only mode when set to yes .
* `mounted` - (Optional) Type: `bool`.
* `newfileman` - (Optional) Type: `string`.
* `nfs_sharing` - (Optional) Type: `bool`.
* `nvme` - (Optional) Type: `string`.
* `nvme_tcp_export` - (Optional) Type: `bool`.
* `nvme_tcp_server_allow_host_name` - (Optional) Type: `string`.
* `nvme_tcp_server_nqn` - (Optional) Type: `string`.
* `nvme_tcp_server_password` - (Optional) Type: `string`.
* `nvme_tcp_server_port` - (Optional) Type: `int`.
* `oldfileman` - (Optional) Type: `string`.
* `parent` - (Optional) Type: `string`. Default: `4.294967295e+09`.
* `part` - (Optional) Type: `string`.
* `partition` - (Optional) Type: `bool`.
* `partition_offset` - (Optional) Type: `string`. Default: `65536`.
* `partition_size` - (Optional) Type: `string`.
* `raid` - (Optional) Type: `string`.
* `raid_and_master` - (Optional) Type: `string`.
* `raid_and_type` - (Optional) Type: `string`.
* `raid_master` - (Optional) Type: `string`. Default: `4.294967295e+09`.
* `raid_member` - (Optional) Type: `bool`.
* `raid_member_failed` - (Optional) Type: `bool`.
* `raid_role` - (Optional) Type: `enum(spare)`.
* `raid_scrub` - (Optional) Type: `string`.
* `read_only` - (Optional) Type: `bool`.
* `reset_counters` - (Optional) Type: `string`.
* `rose` - (Optional) Type: `string`.
* `scan` - (Optional) Type: `string`.
* `self_encrypted_and_locked` - (Optional) Type: `bool`.
* `self_encryption_enabled` - (Optional) Type: `bool`.
* `self_encryption_password` - (Optional) Type: `string`.
* `self_encryption_supported` - (Optional) Type: `bool`.
* `slot` - (Optional) Type: `string`.
* `smb_server_encryption` - (Optional) Type: `bool`.
* `smb_server_password` - (Optional) Type: `string`. **Sensitive.**
* `smb_server_user` - (Optional) Type: `string`.
* `smb_sharing` - (Optional) Type: `bool`.
* `swap` - (Optional) Type: `bool`.
* `swap_enabled` - (Optional) Type: `bool`.
* `tmpfs` - (Optional) Type: `string`.
* `tmpfs_max_size` - (Optional) Type: `string`.
* `trim` - (Optional) Type: `string`.
* `type` - (Optional) Type: `string`. Default: `6`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`. **Marked sensitive**: this menu holds a secret, which RouterOS returns in the row like any other column, so an unprojected read puts it in your state file. Use `proplist` to name the columns you need.
* `active_time` - Type: `string`.
* `available_spare` - Type: `int`.
* `available_spare_threshold` - Type: `int`.
* `controller_burst_time` - Type: `duration`.
* `critical_temperature` - Type: `int`.
* `critical_temperature_time` - Type: `duration`.
* `critical_warning` - Type: `string`.
* `default_slot` - Type: `string`.
* `discard_bytes` - Type: `string`.
* `discard_merges` - Type: `string`.
* `discard_ops` - Type: `string`.
* `discard_time` - Type: `string`.
* `flush_ops` - Type: `string`.
* `flush_time` - Type: `string`.
* `free` - Type: `string`.
* `fs` - Type: `enum(fat32|ext4|btrfs|nfs|smb|wipe, ...)`.
* `fw_version` - Type: `string`.
* `host_read_bytes` - Type: `string`.
* `host_read_commands` - Type: `string`.
* `host_write_bytes` - Type: `string`.
* `host_write_commands` - Type: `string`.
* `in_flight_ops` - Type: `string`.
* `interface` - Type: `string`.
* `interface_speed` - Type: `string`.
* `label` - Type: `string`.
* `model` - Type: `string`.
* `mount_point` - Type: `string`.
* `nvme_tcp_server_secret` - Type: `string`.
* `partition_number` - Type: `int`.
* `percentage_used` - Type: `int`.
* `power_cycles` - Type: `int`.
* `power_on_time` - Type: `duration`.
* `read_bytes` - Type: `string`.
* `read_merges` - Type: `string`.
* `read_ops` - Type: `string`.
* `read_ops_per_second` - Type: `string`.
* `read_rate` - Type: `string`.
* `read_time` - Type: `string`.
* `serial` - Type: `string`.
* `size` - Type: `string`.
* `state` - Type: `string`.
* `temperature` - Type: `int`.
* `temperatures` - Type: `string`.
* `unrecovered_integrity_errors` - Type: `int`.
* `unsafe_shutdown` - Type: `int`.
* `use` - Type: `int`.
* `uuid` - Type: `string`.
* `wait_time` - Type: `string`.
* `warning_temperature` - Type: `int`.
* `warning_temperature_time` - Type: `duration`.
* `write_bytes` - Type: `string`.
* `write_merges` - Type: `string`.
* `write_ops` - Type: `string`.
* `write_ops_per_second` - Type: `string`.
* `write_rate` - Type: `string`.
* `write_time` - Type: `string`.

