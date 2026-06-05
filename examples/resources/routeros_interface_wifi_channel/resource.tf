resource "routeros_interface_wifi_channel" "channel_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # band = "5ghz-ax"
  # channel_width = "replace-me"
  # deprioritize_unii_3_4 = "replace-me"
  # frequency = "2.412e+06"
  # reselect_interval = "replace-me"
  # reselect_time = "replace-me"
  # secondary_frequency = "replace-me"
  # skip_dfs_channels = "replace-me"
}
