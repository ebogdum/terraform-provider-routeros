resource "routeros_interface_wifi_datapath" "datapath_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # bridge = "bridge1"
  # vlan_id = "replace-me"
}
