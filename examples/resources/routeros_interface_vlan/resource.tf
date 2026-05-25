resource "routeros_interface_vlan" "vlan_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  name      = "example"
  vlan_id   = "100"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # l3_hw_offloading = "replace-me"
  # mtu = "replace-me"
  # mvrp = "replace-me"
  # use_service_tag = "replace-me"
}
