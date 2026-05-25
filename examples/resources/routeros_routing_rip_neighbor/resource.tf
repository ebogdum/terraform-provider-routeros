resource "routeros_routing_rip_neighbor" "neighbor_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "10.99.0.1"
  # instance = "replace-me"
}
