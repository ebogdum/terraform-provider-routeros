resource "routeros_ip_proxy" "proxy_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # always_from_cache = false
  # anonymous = false
  # cache_administrator = "replace-me"
  # cache_hit_dscp = 0
  # cache_on_disk = false
  # cache_path = "replace-me"
  # enabled = false
  # max_cache_object_size = 0
  # max_cache_size = "replace-me"
  # max_client_connections = 0
  # max_fresh_time = "1h"
  # max_server_connections = 0
  # parent_proxy = "10.99.0.1"
  # parent_proxy_port = "443"
  # port = "443"
  # serialize_connections = false
  # src_address = "10.99.0.0/24"
}
