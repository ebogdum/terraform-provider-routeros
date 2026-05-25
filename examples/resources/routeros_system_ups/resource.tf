resource "routeros_system_ups" "ups_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # alarm_setting = "immediate"
  # check_capabilities = "replace-me"
  # min_runtime = "replace-me"
  # name = "tf-example"
  # offline_time = "replace-me"
  # port = "443"
}
