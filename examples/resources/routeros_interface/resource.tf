resource "routeros_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # add = "replace-me"
  # edit = "replace-me"
  # find = "replace-me"
  # move = "replace-me"
  # mtu = 1500
  # name = "tf-example"
  # print = "replace-me"
  # remove = "replace-me"
  # set = "replace-me"
}
