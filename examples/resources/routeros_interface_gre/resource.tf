resource "routeros_interface_gre" "gre_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # allow_fast_path = true
  # clamp_tcp_mss = true
  # dont_fragment = "no"
  # dscp = "inherit"
  # ipsec_secret = "REDACTED"
  # keepalive = "1"
  # local_address = "10.99.0.1"
  # mtu = 0
  # name = "tf-example"
  # remote_address = "10.99.0.1"
}
