resource "routeros_certificate_crl" "crl_example" {
  # router = "my-router"  # which router to target; omit for the default
  url = "https://example.com"
}
