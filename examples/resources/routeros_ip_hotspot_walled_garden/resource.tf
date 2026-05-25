resource "routeros_ip_hotspot_walled_garden" "walled_garden_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "replace-me"
  # dst_port = "443"
  # path = "replace-me"
  # server = "replace-me"
  # src_address = "10.99.0.0/24"
}
