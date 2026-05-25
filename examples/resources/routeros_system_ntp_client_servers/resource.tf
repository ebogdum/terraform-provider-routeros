resource "routeros_system_ntp_client_servers" "servers_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "replace-me"
  # auth_key = "REDACTED"
  # iburst = true
  # max_poll = 10
  # min_poll = 6
}
