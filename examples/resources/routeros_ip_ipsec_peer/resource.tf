resource "routeros_ip_ipsec_peer" "peer_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "10.99.0.1"
  # exchange_mode = "aggressive"
  # local_address = "10.99.0.1"
  # name = "example"
  # passive = false
  # port = "443"
  # profile = "replace-me"
  # send_initial_contact = true
}
