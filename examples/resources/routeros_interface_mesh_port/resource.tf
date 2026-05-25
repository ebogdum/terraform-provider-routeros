resource "routeros_interface_mesh_port" "port_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # hello_interval = 10
  # interface = "ether1"
  # mesh = "replace-me"
  # path_cost = 10
  # port_type = "auto"
}
