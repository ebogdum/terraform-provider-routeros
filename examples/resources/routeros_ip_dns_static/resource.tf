resource "routeros_ip_dns_static" "static_example" {
  # router = "my-router"  # which router to target; omit for the default
  address = "127.0.0.1"
  name    = "tf-example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # ttl = "replace-me"
}
