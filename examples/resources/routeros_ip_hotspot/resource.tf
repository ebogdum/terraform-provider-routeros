resource "routeros_ip_hotspot" "hotspot_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  name      = "tf-example"

  disabled = false

  # Optional attributes (uncomment as needed):
  # keepalive_timeout = "replace-me"
  # profile = "replace-me"
}
