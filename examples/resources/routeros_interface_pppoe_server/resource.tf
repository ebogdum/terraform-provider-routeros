resource "routeros_interface_pppoe_server" "pppoe_server_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # name = "tf-example"
  # service = "replace-me"
  # user = "myuser"
}
