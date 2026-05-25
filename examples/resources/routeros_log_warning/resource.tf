resource "routeros_log_warning" "warning_example" {
  # router = "my-router"  # which router to target; omit for the default
  message = "hello from terraform"
}
