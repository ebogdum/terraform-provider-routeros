resource "routeros_interface_dot1x_client" "client_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # anon_identity = "replace-me"
  # certificate = "replace-me"
  # eap_methods = "replace-me"
  # identity = "replace-me"
  # interface = "ether1"
  # password = "REDACTED"
}
