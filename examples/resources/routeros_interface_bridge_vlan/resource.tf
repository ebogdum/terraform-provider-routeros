resource "routeros_interface_bridge_vlan" "vlan_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # bridge = "bridge1"
  # mvrp_attributes = "replace-me"
  # mvrp_forbidden = "replace-me"
  # tagged = "replace-me"
  # untagged = "replace-me"
  # vlan_ids = "replace-me"
}
