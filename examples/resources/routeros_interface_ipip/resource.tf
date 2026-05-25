resource "routeros_interface_ipip" "ipip_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # clamp_tcp_mss = true
  # dont_fragment = "no"
  # dscp = "inherit"
  # ipsec_secret = "REDACTED"
  # keepalive = "replace-me"
  # local_address = "10.99.0.1"
  # mtu = 0
  # name = "tf-example"
  # remote_address = "10.99.0.1"
}
