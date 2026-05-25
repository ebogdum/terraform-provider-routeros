resource "routeros_ipv6_firewall_address_list" "address_list_example" {
  # router = "my-router"  # which router to target; omit for the default
  address = "fd00:db8::/64"
  list    = "my-list"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # parent = 0
  # timeout = "replace-me"
}
