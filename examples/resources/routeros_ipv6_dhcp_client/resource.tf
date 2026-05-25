resource "routeros_ipv6_dhcp_client" "dhcp_client_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  request   = "address"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # clientid = "replace-me"
  # default_route_distance = "replace-me"
  # hostname = "replace-me"
}
