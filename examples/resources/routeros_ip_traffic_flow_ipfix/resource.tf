resource "routeros_ip_traffic_flow_ipfix" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # bytes = true
  # dst_address = true
  # dst_address_mask = true
  # dst_mac_address = true
  # dst_port = true
  # first_forwarded = true
  # gateway = true
  # icmp_code = true
}
