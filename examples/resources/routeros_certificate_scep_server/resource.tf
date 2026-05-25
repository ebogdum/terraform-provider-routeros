resource "routeros_certificate_scep_server" "scep_server_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # ca_certificate = "replace-me"
  # days_valid = 1
  # next_ca_certificate = "replace-me"
  # path = "replace-me"
  # request_lifetime = "3600"
}
