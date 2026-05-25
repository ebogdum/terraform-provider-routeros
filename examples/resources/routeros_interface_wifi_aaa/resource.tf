resource "routeros_interface_wifi_aaa" "aaa_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # called_format = "replace-me"
  # calling_format = "replace-me"
  # interim_update = "replace-me"
  # mac_caching = "replace-me"
  # name = "tf-example"
  # nas_identifier = "replace-me"
  # password_format = "replace-me"
  # username_format = "replace-me"
}
