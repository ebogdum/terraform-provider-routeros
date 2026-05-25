resource "routeros_ip_socksify" "socksify_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # port = "443"
}
