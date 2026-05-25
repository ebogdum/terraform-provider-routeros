resource "routeros_interface_wireguard" "wireguard_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # listen_port = "443"
  # mtu = "replace-me"
  # name = "tf-example"
  # private_key = "REDACTED"
}
