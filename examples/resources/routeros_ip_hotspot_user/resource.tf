resource "routeros_ip_hotspot_user" "user_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "10.99.0.1"
  # def = false
  # email = "replace-me"
  # limit_bytes_in = "replace-me"
  # limit_bytes_out = "replace-me"
  # limit_bytes_total = "replace-me"
  # limit_uptime = "1h"
  # mac_address = "10.99.0.0/24"
  # nondef = "replace-me"
  # nondefro = "replace-me"
  # otp_secret = "REDACTED"
  # password = "REDACTED"
  # profile = "replace-me"
  # reset_all_counters = "replace-me"
  # reset_counters = "replace-me"
  # routes = "replace-me"
  # server = "replace-me"
}
