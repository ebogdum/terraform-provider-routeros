resource "routeros_ip_ipsec_settings" "settings_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # accounting = false
  # ddos_cookie_threshold = 0
  # interim_update = "1h"
  # xauth_use_radius = false
}
