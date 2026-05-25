resource "routeros_interface_macvlan" "macvlan_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # mac_address = "10.99.0.0/24"
  # mode = "replace-me"
  # mtu = "replace-me"
  # name = "tf-example"
}
