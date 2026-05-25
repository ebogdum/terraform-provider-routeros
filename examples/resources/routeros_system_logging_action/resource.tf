resource "routeros_system_logging_action" "action_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # disk_file_count = 0
  # disk_file_name = "replace-me"
  # disk_lines_per_file = 0
  # disk_stop_on_full = false
  # memory_lines = 0
  # memory_stop_on_full = false
  # name = "tf-example"
  # remember = false
  # remote = "replace-me"
  # remote_log_format = "replace-me"
  # remote_port = "443"
  # remote_protocol = "replace-me"
  # src_address = "10.99.0.0/24"
  # target = "replace-me"
  # vrf = "main"
}
