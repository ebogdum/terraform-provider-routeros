resource "routeros_system_scheduler" "scheduler_example" {
  # router = "my-router"  # which router to target; omit for the default
  name     = "tf-example"
  on_event = ":put \"tick\""

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # interval = "1h"
  # policy = "replace-me"
  # start_date = "replace-me"
  # start_time = "startup"
}
