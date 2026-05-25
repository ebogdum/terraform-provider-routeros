resource "routeros_interface_pptp_server" "pptp_server_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"
  user = "myuser"

  comment  = "managed by terraform"
  disabled = false
}
