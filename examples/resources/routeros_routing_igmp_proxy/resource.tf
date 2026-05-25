resource "routeros_routing_igmp_proxy" "igmp_proxy_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # query_interval = "1h"
  # query_response_interval = "1h"
  # quick_leave = false
}
