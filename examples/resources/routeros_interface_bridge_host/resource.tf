resource "routeros_interface_bridge_host" "host_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # bridge = "bridge1"
  # interface = "ether1"
  # mac_address = "10.99.0.0/24"
  # vid = "replace-me"
}
