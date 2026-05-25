resource "routeros_ip_dhcp_server_alert" "alert_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # alert_timeout = "3600"
  # on_alert = "replace-me"
}
