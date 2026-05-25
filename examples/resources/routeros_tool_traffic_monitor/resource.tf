resource "routeros_tool_traffic_monitor" "traffic_monitor_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  name      = "tf-example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # on_event = "replace-me"
  # threshold = "replace-me"
  # traffic = "transmitted"
  # trigger = "v1"
}
