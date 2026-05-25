resource "routeros_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # fp_tx_rx_packet_rate = "replace-me"
  # fp_tx_rx_rate = "replace-me"
  # inactive = false
  # mtu = 1500
  # name = "tf-example"
  # nodefname = "replace-me"
  # notrunning = "replace-me"
  # passthrough = false
  # reset_traffic_counters = "replace-me"
  # slave = false
  # torch = "replace-me"
  # tx_rx_bytes = "replace-me"
  # tx_rx_drops = "replace-me"
  # tx_rx_errors = "replace-me"
  # tx_rx_packet_rate = "replace-me"
  # tx_rx_packets = "replace-me"
  # tx_rx_rate = "replace-me"
}
