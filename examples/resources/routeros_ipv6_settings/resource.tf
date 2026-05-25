resource "routeros_ipv6_settings" "settings_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # accept_redirects = "replace-me"
  # accept_router_advertisements = "replace-me"
  # accept_router_advertisements_on = "replace-me"
  # allow_fast_path = false
  # disable_ipv6 = false
  # disable_link_local_address = "10.99.0.0/24"
  # forward = false
  # max_neighbor_entries = 0
  # min_neighbor_entries = 0
  # multipath_hash_policy = "replace-me"
  # soft_max_neighbor_entries = 0
  # stale_neighbor_detect_interval = 0
  # stale_neighbor_timeout = 0
}
