resource "routeros_disk" "disk_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
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
