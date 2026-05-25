resource "routeros_ip_dhcp_server" "dhcp_server_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  name      = "tf-example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # add_arp_for_leases = false
  # address_list = "replace-me"
  # address_pool = "replace-me"
  # allow_dual_stack_queue = true
  # always_broadcast = false
  # authoritative = "yes"
  # bootp_lease_time = "4.294967295e+09"
  # bootp_support = "static"
  # client_mac_limit = 0
  # conflict_detection = true
  # delay_threshold = "1h"
  # dhcp_option_set = "4.294967295e+09"
  # dynamic_lease_identifiers = "mac-address"
  # dynbootp = "replace-me"
  # insert_queue_before = "0"
  # lease_script = "replace-me"
  # lease_time = "1800"
  # parent_queue = "replace-me"
  # relay = "10.99.0.1"
  # server_address = "10.99.0.0/24"
  # support_the_broadband_forum_tr_101 = false
  # use_framed_as_classless = true
  # use_radius = "no"
  # use_reconfigure = false
}
