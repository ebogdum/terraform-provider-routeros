resource "routeros_routing_bgp_connection" "connection_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # afi = "replace-me"
  # as = "replace-me"
  # connect = "replace-me"
  # hold_time = "replace-me"
  # instance = "replace-me"
  # keepalive_time = "replace-me"
  # listen = "replace-me"
  # multihop = "replace-me"
  # name = "example"
  # nexthop_choice = "replace-me"
  # routing_table = "main"
  # tcp_md5_key = "REDACTED"
  # use_bfd = "replace-me"
  # vrf = "main"
}
