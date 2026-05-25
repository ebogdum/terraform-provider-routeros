resource "routeros_ip_route" "route_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # distance = 0
  # dst_address = "10.99.0.0/24"
  # ecmp = false
  # gateway = "10.99.0.1"
  # hw_offloaded = false
  # routing_table = "main"
  # scope = 0
  # target_scope = 0
  # vrf_interface = "replace-me"
}
