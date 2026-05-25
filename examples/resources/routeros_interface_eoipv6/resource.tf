resource "routeros_interface_eoipv6" "eoipv6_example" {
  # router = "my-router"  # which router to target; omit for the default
  name           = "example"
  remote_address = "10.99.0.1"
  tunnel_id      = "1"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # ipsec_secret = "REDACTED"
  # local_address = "10.99.0.1"
  # mac_address = "10.99.0.0/24"
  # mtu = "replace-me"
}
