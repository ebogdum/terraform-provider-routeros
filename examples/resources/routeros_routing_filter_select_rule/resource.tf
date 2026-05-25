resource "routeros_routing_filter_select_rule" "select_rule_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # chain = "replace-me"
  # type = "where"
}
