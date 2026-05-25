resource "routeros_ip_dhcp_client_option" "option_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # code = 0
  # name = "example"
  # value = "replace-me"
}
