resource "routeros_ip_dhcp_client" "dhcp_client_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # add_default_route = "yes"
  # allow_reconfigure = false
  # check_gateway = "none"
  # default_route_distance = 0
  # default_route_tables = "replace-me"
  # dhcp_options = []
  # dscp = 0
  # interface = "ether1"
  # name = "tf-example"
  # script = "replace-me"
  # use_broadcast = "both"
  # use_peer_dns = true
  # use_peer_ntp = true
  # vlan_priority = 0
}
