resource "routeros_ip_reverse_proxy" "reverse_proxy_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false
}
