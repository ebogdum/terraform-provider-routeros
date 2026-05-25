resource "routeros_ip_dhcp_server_network" "network_example" {
  # router = "my-router"  # which router to target; omit for the default
  address = "10.255.255.0/30"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # boot_file_name = "replace-me"
  # caps_manager = "replace-me"
  # caps_managers = "replace-me"
  # dhcp_option = "replace-me"
  # dhcp_option_set = "4.294967295e+09"
  # dhcp_options = "replace-me"
  # dns_server = "1.1.1.1,8.8.8.8"
  # dns_servers = "replace-me"
  # domain = "example.local"
  # dynamic = "replace-me"
  # gateway = "10.255.255.1"
  # netmask = "255.255.255.0"
  # next_server = "replace-me"
  # nndns = "replace-me"
  # nnntp = "replace-me"
  # no_dns = false
  # no_ntp = false
  # ntp_server = "pool.ntp.org"
  # ntp_servers = "replace-me"
  # wins_server = "replace-me"
  # wins_servers = "replace-me"
}
