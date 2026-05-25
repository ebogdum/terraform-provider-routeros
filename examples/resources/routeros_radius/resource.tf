resource "routeros_radius" "radius_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # accounting_backup = false
  # accounting_port = "443"
  # address = "replace-me"
  # authentication_port = "443"
  # called_id = "replace-me"
  # certificate = "replace-me"
  # domain = "example.local"
  # protocol = "udp"
  # radsec = "replace-me"
  # radsec_timeout = "replace-me"
  # realm = "replace-me"
  # require_message_auth = ""
  # reset_status = "replace-me"
  # secret = "REDACTED"
  # service = "ppp"
  # src_address = "10.99.0.0/24"
  # timeout = 0
  # udp = "replace-me"
}
