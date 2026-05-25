resource "routeros_interface_wifi_access_list" "access_list_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "replace-me"
  # allow_signal_out_of_range = "replace-me"
  # client_isolation = "replace-me"
  # interface = "ether1"
  # mac_address = "10.99.0.0/24"
  # mac_address_mask = "replace-me"
  # multi_passphrase_group = "replace-me"
  # passphrase = "replace-me"
  # radius_accounting = "replace-me"
  # signal_range = "replace-me"
  # ssid_regexp = "replace-me"
  # time = "replace-me"
  # vlan_id = "replace-me"
  # weekdays = "replace-me"
}
