resource "routeros_interface_bridge_filter" "filter_example" {
  # router = "my-router"  # which router to target; omit for the default
  chain = "forward"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "replace-me"
  # dst_address = "10.99.0.0/24"
  # dst_port = "443"
  # in_interface = "ether1"
  # in_interface_list = "LAN"
  # ingress_priority = "replace-me"
  # jump_target = "replace-me"
  # limit = "replace-me"
  # log = "replace-me"
  # log_prefix = "replace-me"
  # out_interface = "ether1"
  # out_interface_list = "LAN"
  # packet_mark = "replace-me"
  # src_address = "10.99.0.0/24"
  # src_mac_address = "10.99.0.0/24"
  # src_port = "443"
  # tls_host = "replace-me"
}
