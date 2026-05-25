resource "routeros_ip_dhcp_server_lease" "lease_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "10.99.0.1"
  # agent_circuit_id = "replace-me"
  # agent_remote_id = "replace-me"
  # allow_dual_stack_queue = "replace-me"
  # always_broadcast = false
  # block_access = false
  # client_id = "replace-me"
  # dhcp_option_set = "4.294967295e+09"
  # insert_queue_before = "replace-me"
  # lease_time = "1h"
  # mac_address = "10.99.0.0/24"
  # parent_queue = "replace-me"
  # queue_type = "replace-me"
  # rate_limit = "replace-me"
  # routes = "replace-me"
  # server = "replace-me"
}
