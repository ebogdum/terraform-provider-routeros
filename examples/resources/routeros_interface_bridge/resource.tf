resource "routeros_interface_bridge" "bridge_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # admin_mac = "replace-me"
  # ageing_time = "1h"
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # auto_mac = false
  # dhcp_snooping = false
  # ether_type = "replace-me"
  # fast_forward = false
  # forward_delay = "1h"
  # frame_types = "replace-me"
  # igmp_snooping = false
  # ingress_filtering = "replace-me"
  # max_learned_entries = "replace-me"
  # max_message_age = "1h"
  # mlag_heartbeat = "1h"
  # mlag_peer_port = "443"
  # mlag_priority = 0
  # mtu = "replace-me"
  # mvrp = "replace-me"
  # name = "tf-example"
  # port_cost_mode = "replace-me"
  # priority = 0
  # protocol_mode = "replace-me"
  # pvid = "replace-me"
  # ra_guard = false
  # region_name = "replace-me"
  # region_revision = "replace-me"
  # transmit_hold_count = 0
  # vlan_filtering = false
}
