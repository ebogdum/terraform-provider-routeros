resource "routeros_interface_pptp_client" "pptp_client_example" {
  # router = "my-router"  # which router to target; omit for the default
  connect_to = "127.0.0.1"
  name       = "tf-example"
  user       = "myuser"

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
  # password = "REDACTED"
  # profile = "replace-me"
}
