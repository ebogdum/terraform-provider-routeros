resource "routeros_system_console" "console_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # channel = 0
  # port = "443"
  # term = "replace-me"
}
