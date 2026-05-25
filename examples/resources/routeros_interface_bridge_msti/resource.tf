resource "routeros_interface_bridge_msti" "msti_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # bridge = "bridge1"
  # identifier = 0
  # priority = 32768
  # status = 0
  # vlan_mapping = "replace-me"
}
