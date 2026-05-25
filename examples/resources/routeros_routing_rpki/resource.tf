resource "routeros_routing_rpki" "rpki_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "replace-me"
  # expire_interval = "replace-me"
  # group = "replace-me"
  # port = "443"
  # preference = "replace-me"
  # refresh_interval = "replace-me"
  # retry_interval = "replace-me"
  # vrf = "main"
}
