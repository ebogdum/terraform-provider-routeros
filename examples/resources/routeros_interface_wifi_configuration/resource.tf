resource "routeros_interface_wifi_configuration" "configuration_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # antenna_gain = "replace-me"
  # channel = "replace-me"
  # country = "replace-me"
  # distance = "replace-me"
  # mode = "replace-me"
  # ssid = "replace-me"
  # tx_power = "replace-me"
}
