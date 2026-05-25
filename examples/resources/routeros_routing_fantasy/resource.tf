resource "routeros_routing_fantasy" "fantasy_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # dst_address = "10.99.0.0/24"
  # gateway = "replace-me"
  # name = "example"
  # scope = "replace-me"
  # target_scope = "replace-me"
}
