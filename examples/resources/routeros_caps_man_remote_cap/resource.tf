resource "routeros_caps_man_remote_cap" "remote_cap_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false
}
