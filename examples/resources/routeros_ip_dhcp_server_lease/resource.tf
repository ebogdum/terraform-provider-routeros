resource "routeros_ip_dhcp_server_lease" "lease_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "replace-me"
  # address_list = "replace-me"
  # agent_circuit_id = "replace-me"
  # agent_remote_id = "replace-me"
  # allow_dual_stack_queue = "replace-me"
  # always_broadcast = false
  # block_access = false
  # blocked = false
  # check_status = "replace-me"
  # client_id = "replace-me"
  # dhcp_option_set = "4.294967295e+09"
  # dhcp_options = "replace-me"
  # dyn = "replace-me"
  # insert_queue_before = "replace-me"
  # lease_time = "1h"
  # mac_address = "10.99.0.0/24"
  # make_static = "replace-me"
  # parent_queue = "replace-me"
  # ping = "replace-me"
  # queue_type = "replace-me"
  # radius = false
  # rate_limit = "replace-me"
  # rostat = "replace-me"
  # routes = "replace-me"
  # send_reconfigure = "replace-me"
  # server = "replace-me"
  # stat = "replace-me"
  # use_src_mac_address = "10.99.0.0/24"
}
