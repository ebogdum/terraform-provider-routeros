resource "routeros_ppp_secret" "secret_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # caller_id = "replace-me"
  # ipv6 = "replace-me"
  # ipv6_routes = "replace-me"
  # limit_bytes_in = "replace-me"
  # limit_bytes_out = "replace-me"
  # local_address = "10.99.0.1"
  # password = "REDACTED"
  # profile = "replace-me"
  # remote_address = "10.99.0.1"
  # remote_ipv6_prefix = "replace-me"
  # routes = "replace-me"
  # service = "any"
}
