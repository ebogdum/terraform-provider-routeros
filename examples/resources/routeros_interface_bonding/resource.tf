resource "routeros_interface_bonding" "bonding_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # mode = "replace-me"
  # mtu = "replace-me"
  # name = "example"
}
