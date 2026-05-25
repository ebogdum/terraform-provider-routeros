resource "routeros_ip_dns_static" "static_example" {
  # router = "my-router"  # which router to target; omit for the default
  address = "127.0.0.1"
  name    = "tf-example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address_list = "replace-me"
  # cname = "replace-me"
  # forward_to = "replace-me"
  # match_subdomain = "replace-me"
  # mx_exchange = "replace-me"
  # mx_preference = "replace-me"
  # ns = "replace-me"
  # regexp = "replace-me"
  # srv_port = "443"
  # srv_priority = "replace-me"
  # srv_target = "replace-me"
  # srv_weight = "replace-me"
  # text = "replace-me"
  # ttl = "replace-me"
  # type = "replace-me"
}
