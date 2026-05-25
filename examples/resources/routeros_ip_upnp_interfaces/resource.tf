resource "routeros_ip_upnp_interfaces" "interfaces_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  type      = "internal"

  disabled = false
}
