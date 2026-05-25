resource "routeros_ppp_l2tp_secret" "l2tp_secret_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # address = "10.99.0.1"
  # secret = "REDACTED"
}
