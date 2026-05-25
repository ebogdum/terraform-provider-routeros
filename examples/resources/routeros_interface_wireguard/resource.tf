resource "routeros_interface_wireguard" "wireguard_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # l2_mtu = 1500
  # listen_port = "443"
  # mtu = 1420
  # name = "tf-example"
  # private_key = "REDACTED"
  # wg_export = "replace-me"
}
