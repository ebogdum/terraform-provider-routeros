resource "routeros_interface_wifi_datapath" "datapath_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # bridge = "bridge1"
  # bridge_cost = "replace-me"
  # bridge_horizon = "replace-me"
  # client_isolation = "replace-me"
  # interface_list = "replace-me"
  # open_flow_switch = "replace-me"
  # openflow = "replace-me"
  # traffic_processing = "replace-me"
  # vlan_id = "replace-me"
}
