resource "routeros_ipv6_neighbor" "neighbor_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "10.99.0.1"
  # interface = "ether1"
  # mac_address = "10.99.0.0/24"
}
