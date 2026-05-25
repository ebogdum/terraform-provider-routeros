resource "routeros_interface_list_member" "member_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # dynamic = "replace-me"
  # interface = "ether1"
  # list = "my-list"
}
