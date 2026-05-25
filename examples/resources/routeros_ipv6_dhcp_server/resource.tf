resource "routeros_ipv6_dhcp_server" "dhcp_server_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  name      = "tf-example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # dhcp_option = "replace-me"
  # lease_time = "replace-me"
}
