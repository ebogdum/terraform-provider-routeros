resource "routeros_ip_dhcp_server_config" "config_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # accounting = false
  # interim_update = "1h"
  # radius_password = "REDACTED"
  # store_leases_disk = "1h"
}
