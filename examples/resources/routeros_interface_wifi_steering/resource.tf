resource "routeros_interface_wifi_steering" "steering_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "example"

  comment  = "managed by terraform"
  disabled = false
}
