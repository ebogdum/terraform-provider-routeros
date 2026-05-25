resource "routeros_routing_bgp_vpn" "vpn_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # instance = "replace-me"
  # label_allocation_policy = ""
  # route_distinguisher = "replace-me"
  # vrf = "main"
}
