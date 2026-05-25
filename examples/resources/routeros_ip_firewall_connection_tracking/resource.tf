resource "routeros_ip_firewall_connection_tracking" "tracking_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # enabled = "replace-me"
  # generic_timeout = "1h"
  # icmp_timeout = "1h"
  # liberal_tcp_tracking = false
  # loose_tcp_tracking = false
  # tcp_close_timeout = "1h"
  # tcp_close_wait_timeout = "1h"
  # tcp_established_timeout = "1h"
  # tcp_fin_wait_timeout = "1h"
  # tcp_last_ack_timeout = "1h"
  # tcp_max_retrans_timeout = "1h"
  # tcp_syn_received_timeout = "1h"
  # tcp_syn_sent_timeout = "1h"
  # tcp_time_wait_timeout = "1h"
  # tcp_unacked_timeout = "1h"
  # udp_stream_timeout = "1h"
  # udp_timeout = "1h"
}
