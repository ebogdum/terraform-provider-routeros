resource "routeros_ip_address" "address_example" {
  # router = "my-router"  # which router to target; omit for the default
  address   = "10.99.0.1"
  interface = "ether1"

  comment  = "managed by terraform"
  disabled = false
}
