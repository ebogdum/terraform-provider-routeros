resource "routeros_ip_settings" "settings_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # accept_redirects = false
  # accept_source_route = false
  # allow_fast_path = false
  # arp_timeout = "1h"
  # icmp_errors_use_inbound_interface_address = "10.99.0.0/24"
  # icmp_rate_limit = 0
  # icmp_rate_mask = 0
  # ip_forward = false
  # ipv4_multipath_hash_policy = "replace-me"
  # max_neighbor_entries = 0
  # rp_filter = false
  # secure_redirects = false
  # send_redirects = false
  # tcp_syncookies = false
  # tcp_timestamps = "replace-me"
}
