resource "routeros_system_ntp_server" "server_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # auth_key = "REDACTED"
  # broadcast = false
  # broadcast_addresses = "replace-me"
  # enabled = false
  # local_clock_stratum = 0
  # manycast = false
  # multicast = false
  # use_local_clock = false
  # vrf = "main"
}
