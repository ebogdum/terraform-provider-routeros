resource "routeros_interface_ethernet" "ethernet_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "10.99.0.1"
  # advertise = []
  # arp = "disabled"
  # arp_timeout = "1h"
  # cable_settings = "replace-me"
  # combo_mode = "auto"
  # disable_running_check = false
  # fec_mode = "off"
  # interface = "ether1"
  # l2mtu = "replace-me"
  # loop_protect = "default"
  # loop_protect_disable_time = "1h"
  # loop_protect_send_interval = "1h"
  # mac_address = "10.99.0.0/24"
  # mtu = 1500
  # name = "example"
  # orig_mac_address = "10.99.0.0/24"
  # published = "replace-me"
  # rx_flow_control = "off"
  # sfp_shutdown_temperature = 0
  # speed = "10M baseT half"
  # tx_flow_control = "off"
}
