resource "routeros_interface_ovpn_client" "ovpn_client_example" {
  # router = "my-router"  # which router to target; omit for the default
  connect_to = "127.0.0.1"
  name       = "example"
  user       = "myuser"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # certificate = "replace-me"
  # cipher = "replace-me"
  # mac_address = "10.99.0.0/24"
  # max_mtu = "replace-me"
  # mode = "replace-me"
  # password = "REDACTED"
  # profile = "replace-me"
  # verify_server_certificate = "replace-me"
}
