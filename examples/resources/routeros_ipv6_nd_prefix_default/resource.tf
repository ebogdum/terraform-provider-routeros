resource "routeros_ipv6_nd_prefix_default" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # autonomous = true
  # dhcp6_pd_preferred = true
  # preferred_lifetime = "replace-me"
  # valid_lifetime = "replace-me"
}
