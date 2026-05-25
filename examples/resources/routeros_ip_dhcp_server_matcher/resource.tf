resource "routeros_ip_dhcp_server_matcher" "matcher_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false
}
