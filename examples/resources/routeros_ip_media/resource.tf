resource "routeros_ip_media" "media_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"

  disabled = false

  # Optional attributes (uncomment as needed):
  # path = "replace-me"
}
