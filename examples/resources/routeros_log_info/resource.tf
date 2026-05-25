resource "routeros_log_info" "info_example" {
  # router = "my-router"  # which router to target; omit for the default
  message = "hello from terraform"
}
