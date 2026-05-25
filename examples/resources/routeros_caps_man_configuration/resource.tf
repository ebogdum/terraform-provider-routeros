resource "routeros_caps_man_configuration" "configuration_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # channel = "replace-me"
  # country = "replace-me"
  # distance = "replace-me"
  # mode = "replace-me"
  # ssid = "replace-me"
}
