resource "routeros_ip_traffic_flow" "traffic_flow_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # active_flow_timeout = "1h"
  # cache_entries = "replace-me"
  # enabled = false
  # inactive_flow_timeout = "1h"
  # interfaces = "replace-me"
  # packet_sampling = false
  # sampling_interval = 0
  # sampling_space = 0
}
