resource "routeros_ipv6_nd" "nd_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # advertise_dns = "no"
  # advertise_mac_address = "10.99.0.0/24"
  # dns_servers = "replace-me"
  # hop_limit = 64
  # interface = "ether1"
  # managed_address_configuration = false
  # mtu = 0
  # other_configuration = false
  # pref64_prefixes = "replace-me"
  # ra_delay = "3"
  # ra_interval = "replace-me"
  # ra_lifetime = "1800"
  # ra_preference = "medium"
  # reachable_time = 0
  # retransmit_interval = 0
}
