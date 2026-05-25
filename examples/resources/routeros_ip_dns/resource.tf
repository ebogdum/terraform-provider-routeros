resource "routeros_ip_dns" "dns_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # address_list_extra_time = "1h"
  # allow_remote_requests = false
  # cache_max_ttl = "1h"
  # cache_size = 0
  # doh_max_concurrent_queries = 0
  # doh_max_server_connections = 0
  # doh_timeout = "1h"
  # max_concurrent_queries = 0
  # max_concurrent_tcp_sessions = 0
  # max_udp_packet_size = 0
  # mdns_repeat_ifaces = "replace-me"
  # query_server_timeout = "1h"
  # query_total_timeout = "1h"
  # servers = "replace-me"
  # use_doh_server = "replace-me"
  # verify_doh_cert = false
  # vrf = "main"
}
