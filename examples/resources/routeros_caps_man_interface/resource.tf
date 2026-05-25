resource "routeros_caps_man_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp_timeout = "replace-me"
  # mac_address = "10.99.0.0/24"
  # master_interface = "ether1"
  # name = "tf-example"
  # radio_mac = "02:00:00:00:00:01"
}
