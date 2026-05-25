resource "routeros_interface_wifi_channel" "channel_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # band = "replace-me"
  # frequency = "replace-me"
}
