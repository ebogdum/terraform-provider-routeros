resource "routeros_routing_gmp" "gmp_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # exclude = false
  # group = "replace-me"
  # interfaces = "replace-me"
  # sources = "replace-me"
}
