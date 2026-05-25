resource "routeros_interface_bridge_port" "port_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # auto_isolate = false
  # bpdu_guard = false
  # bridge = "bridge1"
  # broadcast_flood = true
  # edge = "auto"
  # fast_leave = false
  # frame_types = "admit all"
  # horizon = 0
  # hw = "replace-me"
  # ingress_filtering = true
  # interface = "ether1"
  # internal_path_cost = "replace-me"
  # learn = "auto"
  # multicast_router = "Disabled"
  # mvrp_applicant_state = "normal participant"
  # mvrp_registrar_state = "normal"
  # path_cost = "replace-me"
  # point_to_point = "auto"
  # priority = 128
  # pvid = 0
  # restricted_role = false
  # restricted_tcn = false
  # tag_stacking = false
  # trusted = false
  # trusted_ra = false
  # unknown_multicast_flood = true
  # unknown_unicast_flood = true
}
