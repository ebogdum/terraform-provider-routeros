resource "routeros_routing_isis_interface_template" "interface_template_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false
}
