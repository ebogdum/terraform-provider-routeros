resource "routeros_interface_macsec" "macsec_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # mtu = "replace-me"
  # name = "example"
  # profile = "replace-me"
}
