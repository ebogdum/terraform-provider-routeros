resource "routeros_ip_firewall_calea" "calea_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false
}
