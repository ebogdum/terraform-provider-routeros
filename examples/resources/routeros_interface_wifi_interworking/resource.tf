resource "routeros_interface_wifi_interworking" "interworking_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment  = "managed by terraform"
  disabled = false
}
