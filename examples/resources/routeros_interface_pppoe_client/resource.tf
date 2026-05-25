resource "routeros_interface_pppoe_client" "pppoe_client_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # ac_name = "replace-me"
  # add_default_route = "replace-me"
  # allow = "replace-me"
  # default_route_distance = "replace-me"
  # dial_on_demand = "replace-me"
  # interface = "ether1"
  # keepalive_timeout = "replace-me"
  # max_mru = "replace-me"
  # max_mtu = "replace-me"
  # mrru = "replace-me"
  # name = "tf-example"
  # password = "REDACTED"
  # profile = "replace-me"
  # service_name = "replace-me"
  # use_peer_dns = "replace-me"
  # user = "myuser"
}
