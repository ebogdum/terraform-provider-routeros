resource "routeros_interface_vxlan" "vxlan_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"
  vni  = "100"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # bridge = "bridge1"
  # interface = "ether1"
  # local_address = "10.99.0.1"
  # mac_address = "10.99.0.0/24"
  # mtu = "replace-me"
  # port = "443"
  # ttl = "replace-me"
}
