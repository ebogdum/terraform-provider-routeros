resource "routeros_tool_traffic_generator_port" "port_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # interface = "ether1"
  # name = "example"
}
