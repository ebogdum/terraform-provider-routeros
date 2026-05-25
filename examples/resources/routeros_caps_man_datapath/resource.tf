resource "routeros_caps_man_datapath" "datapath_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # arp = "replace-me"
  # bridge = "bridge1"
  # l2mtu = "replace-me"
  # mtu = "replace-me"
  # vlan_id = "replace-me"
}
