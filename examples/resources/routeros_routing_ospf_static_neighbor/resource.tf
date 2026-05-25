resource "routeros_routing_ospf_static_neighbor" "static_neighbor_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "replace-me"
  # area = "replace-me"
  # instance_id = 0
  # poll_interval = "1h"
}
