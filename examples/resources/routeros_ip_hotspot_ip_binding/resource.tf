resource "routeros_ip_hotspot_ip_binding" "ip_binding_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "10.99.0.0/24"
  # mac_address = "10.99.0.0/24"
  # server = "replace-me"
  # to_address = "10.99.0.0/24"
  # type = "regular"
}
