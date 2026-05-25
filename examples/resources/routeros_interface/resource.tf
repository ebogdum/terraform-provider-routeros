resource "routeros_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # mtu = 1500
  # name = "tf-example"
}
