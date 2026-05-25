resource "routeros_system_watchdog" "watchdog_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # auto_send_supout = false
  # automatic_supout = false
  # no_ping_delay = "replace-me"
  # ping_start_after_boot = "1h"
  # ping_timeout = "1h"
  # send_email_from = "replace-me"
  # send_email_to = "replace-me"
  # send_smtp_server = "replace-me"
  # watch_address = "10.99.0.0/24"
  # watchdog_timer = false
}
