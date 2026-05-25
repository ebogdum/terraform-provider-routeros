resource "routeros_interface_dot1x_server" "server_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # accounting = true
  # auth_timeout = "6000"
  # auth_types = "1"
  # guest_vlan_id = "replace-me"
  # interface = "ether1"
  # interim_update = "1h"
  # mac = "replace-me"
  # mac_auth_mode = "mac-as-username"
  # radius_mac_format = "xx:xx:xx:xx:xx:xx"
  # reauth_timeout = "replace-me"
  # reject_vlan_id = "replace-me"
  # retrans_timeout = "3000"
  # server_fail_vlan_id = "replace-me"
}
