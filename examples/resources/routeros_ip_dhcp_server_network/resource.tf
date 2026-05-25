resource "routeros_ip_dhcp_server_network" "network_example" {
  # router = "my-router"  # which router to target; omit for the default
  address = "10.99.0.1"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # boot_file_name = "replace-me"
  # dhcp_option_set = "4.294967295e+09"
  # domain = "example.local"
  # gateway = "10.255.255.1"
  # netmask = "255.255.255.0"
  # next_server = "replace-me"
}
