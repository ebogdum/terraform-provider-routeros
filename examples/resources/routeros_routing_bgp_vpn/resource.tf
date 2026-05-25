resource "routeros_routing_bgp_vpn" "vpn_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # export_filter = "replace-me"
  # export_route_targets = "replace-me"
  # export_select = "replace-me"
  # import_filter = "replace-me"
  # import_route_targets = "replace-me"
  # instance = "replace-me"
  # label_allocation_policy = ""
  # redistribute = "replace-me"
  # route_distinguisher = "replace-me"
  # vrf = "main"
}
