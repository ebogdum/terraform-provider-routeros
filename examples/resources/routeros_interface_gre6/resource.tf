resource "routeros_interface_gre6" "gre6_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # ipsec_secret = "REDACTED"
  # local_address = "10.99.0.1"
  # mtu = "replace-me"
  # name = "example"
  # remote_address = "10.99.0.1"
}
