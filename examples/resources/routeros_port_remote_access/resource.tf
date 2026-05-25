resource "routeros_port_remote_access" "remote_access_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # channel = "replace-me"
  # local_address = "10.99.0.1"
  # port = "443"
  # protocol = "replace-me"
}
