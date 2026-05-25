resource "routeros_interface_eoip" "eoip_example" {
  # router = "my-router"  # which router to target; omit for the default
  name           = "tf-example"
  remote_address = "10.99.0.1"
  tunnel_id      = 1

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # allow_fast_path = true
  # arp = "enabled"
  # arp_timeout = "1h"
  # clamp_tcp_mss = true
  # dont_fragment = "no"
  # dscp = "inherit"
  # ipsec_secret = "REDACTED"
  # keepalive = "1"
  # local_address = "10.99.0.1"
  # loop_protect = "default"
  # loop_protect_disable_time = "replace-me"
  # loop_protect_send_interval = "replace-me"
  # mac_address = "10.99.0.0/24"
  # mtu = 0
}
