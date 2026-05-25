resource "routeros_interface_bridge" "bridge_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # ageing_time = "1h"
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # auto_mac = false
  # dhcp_snooping = false
  # fast_forward = false
  # forward_delay = "1h"
  # igmp_snooping = false
  # max_learned_entries = "replace-me"
  # max_message_age = "1h"
  # mlag_heartbeat = "1h"
  # mlag_peer_port = "443"
  # mlag_priority = 0
  # mtu = "replace-me"
  # name = "tf-example"
  # port_cost_mode = "replace-me"
  # priority = 0
  # protocol_mode = "replace-me"
  # ra_guard = false
  # transmit_hold_count = 0
  # vlan_filtering = false
}
