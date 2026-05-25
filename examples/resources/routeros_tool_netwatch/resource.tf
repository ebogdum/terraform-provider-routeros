resource "routeros_tool_netwatch" "netwatch_example" {
  # router = "my-router"  # which router to target; omit for the default
  host = "1.1.1.1"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # certificate = "replace-me"
  # dns_server = "1.1.1.1,8.8.8.8"
  # interval = "1m"
  # name = "tf-example"
  # port = "443"
  # src_address = "10.99.0.0/24"
  # timeout = "replace-me"
  # ttl = "replace-me"
  # type = "icmp"
}
