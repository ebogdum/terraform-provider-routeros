resource "routeros_ipv6_nd_prefix" "prefix_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # x6to4_interface = "4.294967295e+09"
  # autonomous = true
  # dhcpv6_pd_preferred = false
  # interface = "ether1"
  # no6to4 = "replace-me"
  # on_link = true
  # preferred_lifetime = "604800"
  # prefix = "replace-me"
  # valid_lifetime = "2.592e+06"
}
