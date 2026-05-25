resource "routeros_caps_man_radio" "radio_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false
}
