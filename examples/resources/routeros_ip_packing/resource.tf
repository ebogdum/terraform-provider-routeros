resource "routeros_ip_packing" "packing_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"

  disabled = false
}
