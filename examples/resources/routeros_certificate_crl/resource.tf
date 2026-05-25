resource "routeros_certificate_crl" "crl_example" {
  # router = "my-router"  # which router to target; omit for the default
  url = "https://example.com"

  # Optional attributes (uncomment as needed):
  # download = "replace-me"
  # expired = false
  # flush = "replace-me"
}
