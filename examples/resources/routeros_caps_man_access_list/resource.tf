resource "routeros_caps_man_access_list" "access_list_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "replace-me"
  # interface = "ether1"
  # mac_address = "10.99.0.0/24"
  # vlan_id = "replace-me"
}
