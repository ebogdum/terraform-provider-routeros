resource "routeros_mpls_ldp_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"

  comment  = "managed by terraform"
  disabled = false
}
