resource "routeros_interface_bridge_host" "host_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # aged = false
  # aged_on_peer = false
  # bridge = "bridge1"
  # external_fdb = false
  # interface = "ether1"
  # local = false
  # mac_address = "10.99.0.0/24"
  # vid = "replace-me"
}
