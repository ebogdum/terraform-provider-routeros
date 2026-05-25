resource "routeros_ipv6_dhcp_relay" "dhcp_relay_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false
}
