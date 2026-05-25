resource "routeros_interface_ovpn_server" "ovpn_server_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "example"
  user = "myuser"

  comment  = "managed by terraform"
  disabled = false
}
