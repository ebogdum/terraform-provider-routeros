resource "routeros_routing_bgp_template" "template_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # afi = "replace-me"
  # as = "65000"
  # hold_time = "replace-me"
  # keepalive_time = "replace-me"
  # multihop = "replace-me"
  # nexthop_choice = "replace-me"
  # router_id = "1.1.1.1"
  # routing_table = "main"
  # templates = "replace-me"
  # use_bfd = "replace-me"
  # vrf = "main"
}
