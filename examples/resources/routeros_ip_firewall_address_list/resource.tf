resource "routeros_ip_firewall_address_list" "address_list_example" {
  # router = "my-router"  # which router to target; omit for the default
  address = "10.99.0.1"
  list    = "my-list"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # timeout = "replace-me"
}
