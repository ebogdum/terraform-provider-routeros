resource "routeros_ip_arp" "arp_example" {
  # router = "my-router"  # which router to target; omit for the default
  address   = "10.99.0.1"
  interface = "ether1"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # ip_address = "10.99.0.0/24"
  # mac_address = "10.99.0.0/24"
  # mac_ping = "replace-me"
  # mac_telnet = "replace-me"
  # make_static = "replace-me"
  # ping = "replace-me"
  # published = false
  # telnet = "replace-me"
  # torch = "replace-me"
}
