resource "routeros_ip_arp" "arp_example" {
  # router = "my-router"  # which router to target; omit for the default
  address   = "10.99.0.1"
  interface = "ether1"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # mac_address = "10.99.0.0/24"
  # published = false
}
