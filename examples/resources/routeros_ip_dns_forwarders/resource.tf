resource "routeros_ip_dns_forwarders" "forwarders_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # dns_servers = "replace-me"
  # name = "tf-example"
}
