resource "routeros_interface_vpls" "vpls_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "enabled"
  # arp_timeout = "1h"
  # bgp_signaled = false
  # bridge = "bridge1"
  # bridge_cost = "replace-me"
  # bridge_horizon = "replace-me"
  # bridge_pvid = "replace-me"
  # cisco_bgp_signaled = false
  # cisco_static_id = "replace-me"
  # mac_address = "10.99.0.0/24"
  # mtu = 1500
  # pw_control_word = "replace-me"
  # pw_l2mtu = "replace-me"
  # pw_type = "replace-me"
  # remote_peer = "replace-me"
  # vpls_id = "replace-me"
}
