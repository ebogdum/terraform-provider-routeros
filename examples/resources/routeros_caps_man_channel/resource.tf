resource "routeros_caps_man_channel" "channel_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "example"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # band = "replace-me"
  # frequency = "replace-me"
  # tx_power = "replace-me"
}
