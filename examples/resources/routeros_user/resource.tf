resource "routeros_user" "user_example" {
  # router = "my-router"  # which router to target; omit for the default
  group    = "read"
  name     = "example"
  password = "REDACTED"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "10.99.0.1"
  # inactivity_policy = "replace-me"
  # inactivity_timeout = "1h"
}
