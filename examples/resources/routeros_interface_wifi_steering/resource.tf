resource "routeros_interface_wifi_steering" "steering_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # x2g_probe_delay = "replace-me"
  # neighbor_group = "replace-me"
  # neighbor_groups = "replace-me"
  # rrm = "replace-me"
  # transition_request_count = "replace-me"
  # transition_threshold = "replace-me"
  # transition_threshold_period = "replace-me"
  # transition_threshold_time = "replace-me"
  # transition_time = "replace-me"
  # wnm = "replace-me"
}
