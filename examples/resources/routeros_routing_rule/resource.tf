resource "routeros_routing_rule" "rule_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "replace-me"
  # chain = "replace-me"
  # dst_address = "10.99.0.0/24"
  # interface = "ether1"
  # realm = "replace-me"
  # routing_mark = "replace-me"
  # src_address = "10.99.0.0/24"
  # vrf = "main"
}
