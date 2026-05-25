resource "routeros_ip_dhcp_server_option" "option_example" {
  # router = "my-router"  # which router to target; omit for the default
  code  = 60
  name  = "tf-example"
  value = "'tf-acc'"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # force = false
}
