resource "routeros_ip_hotspot_walled_garden_ip" "ip_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "accept"
  # dst_address = "10.99.0.0/24"
  # dst_address_list = "my-list"
  # dst_host = "replace-me"
  # dst_port = "443"
  # protocol = "replace-me"
  # server = "replace-me"
  # src_address = "10.99.0.0/24"
  # src_address_list = "my-list"
}
