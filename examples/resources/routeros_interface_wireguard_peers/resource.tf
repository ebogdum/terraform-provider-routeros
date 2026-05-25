resource "routeros_interface_wireguard_peers" "peers_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # allowed_address = "10.99.0.0/24"
  # client_address = "10.99.0.0/24"
  # client_allowed_address = "10.99.0.0/24"
  # client_dns = "replace-me"
  # client_endpoint = "replace-me"
  # client_keepalive = "1h"
  # client_listen_port = "443"
  # endpoint = "replace-me"
  # endpoint_address = "10.99.0.0/24"
  # endpoint_port = "443"
  # interface = "ether1"
  # name = "tf-example"
  # persistent_keepalive = "1h"
  # preshared_key = "REDACTED"
  # private_key = "REDACTED"
  # public_key = "REDACTED"
  # responder = false
}
