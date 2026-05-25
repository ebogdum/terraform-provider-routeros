resource "routeros_interface_w60g" "w60g_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # frequency = "replace-me"
  # isolate_stations = "replace-me"
  # l2mtu = "replace-me"
  # mac_address = "10.99.0.0/24"
  # mdmg_fix = "replace-me"
  # mode = "replace-me"
  # mtu = "replace-me"
  # name = "tf-example"
  # password = "REDACTED"
  # put_stations_in_bridge = "replace-me"
  # region = "replace-me"
  # scan_list = "replace-me"
  # ssid = "replace-me"
  # tx_sector = "replace-me"
}
