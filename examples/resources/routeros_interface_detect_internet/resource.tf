resource "routeros_interface_detect_internet" "detect_internet_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # detect_interface_list = "replace-me"
  # internet_interface_list = "replace-me"
  # lan_interface_list = "replace-me"
  # request_interval = "1h"
  # wan_interface_list = "replace-me"
}
