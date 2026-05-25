resource "routeros_interface_lte_apn" "apn_example" {
  # router = "my-router"  # which router to target; omit for the default
  apn  = "internet"
  name = "tf-example"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # add_default_route = false
  # authentication = "replace-me"
  # default_route_distance = 0
  # ip_type = "replace-me"
  # use_network_apn = false
  # use_peer_dns = false
}
