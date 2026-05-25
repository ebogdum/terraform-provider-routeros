resource "routeros_ip_dhcp_relay" "dhcp_relay_example" {
  # router = "my-router"  # which router to target; omit for the default
  dhcp_server = "127.0.0.1"
  interface   = "ether1"
  name        = "tf-example"

  disabled = false

  # Optional attributes (uncomment as needed):
  # add_relay_info = false
  # delay_threshold = "1h"
  # dhcp_server_vrf = "replace-me"
  # local_address = "10.99.0.1"
  # local_address_as_source_ip = false
  # relay_info_remote_id = "replace-me"
  # reset_counters = "replace-me"
}
