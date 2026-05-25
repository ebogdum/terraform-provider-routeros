resource "routeros_routing_filter_rule" "rule_example" {
  # router = "my-router"  # which router to target; omit for the default
  chain = "tf_acc_chain"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # rule = "accept"
}
