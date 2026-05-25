resource "routeros_queue_simple" "simple_example" {
  # router = "my-router"  # which router to target; omit for the default
  name   = "example"
  target = "127.0.0.1/32"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # burst_limit = "replace-me"
  # burst_threshold = "replace-me"
  # burst_time = "replace-me"
  # limit_at = "replace-me"
  # max_limit = "replace-me"
  # packet_marks = "replace-me"
  # parent = "replace-me"
  # place_before = "replace-me"
  # priority = "replace-me"
  # queue = "replace-me"
}
