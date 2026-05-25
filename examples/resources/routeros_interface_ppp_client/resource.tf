resource "routeros_interface_ppp_client" "ppp_client_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # allow = "replace-me"
  # default_route_distance = "replace-me"
  # dial_on_demand = "replace-me"
  # keepalive_timeout = "replace-me"
  # max_mru = "replace-me"
  # max_mtu = "replace-me"
  # mrru = "replace-me"
  # name = "example"
  # password = "REDACTED"
  # profile = "replace-me"
  # remote_address = "10.99.0.1"
  # user = "myuser"
}
