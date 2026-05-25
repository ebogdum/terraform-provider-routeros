resource "routeros_interface_ppp_server" "ppp_server_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # authentication = "replace-me"
  # max_mru = "replace-me"
  # max_mtu = "replace-me"
  # mrru = "replace-me"
  # name = "tf-example"
  # profile = "replace-me"
}
