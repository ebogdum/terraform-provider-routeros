resource "routeros_ip_traffic_flow_target" "target_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # dst_address = "10.99.0.0/24"
  # port = "443"
  # src_address = "10.99.0.0/24"
  # version = "9"
}
