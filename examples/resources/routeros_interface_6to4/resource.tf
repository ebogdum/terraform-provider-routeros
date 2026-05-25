resource "routeros_interface_6to4" "6to4_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # clamp_tcp_mss = false
  # dont_fragment = false
  # dscp = "replace-me"
  # local_address = "10.99.0.1"
  # mtu = "replace-me"
  # name = "example"
  # remote_address = "10.99.0.1"
}
