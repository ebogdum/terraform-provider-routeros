resource "routeros_certificate_scep_server" "scep_server_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # days_valid = 1
  # path = "replace-me"
  # request_lifetime = "3600"
}
