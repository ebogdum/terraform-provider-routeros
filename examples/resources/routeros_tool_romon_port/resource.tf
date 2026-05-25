resource "routeros_tool_romon_port" "port_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # cost = 100
  # forbid = false
  # interface = "ether1"
  # secrets = "replace-me"
}
