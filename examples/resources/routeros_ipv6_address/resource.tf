resource "routeros_ipv6_address" "address_example" {
  # router = "my-router"  # which router to target; omit for the default
  address   = "10.99.0.1"
  interface = "ether1"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # advertise = true
  # auto_link_local = false
  # eui_64 = false
  # from_pool = "replace-me"
  # no_dad = false
}
