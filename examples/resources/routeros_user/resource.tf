resource "routeros_user" "user_example" {
  # router = "my-router"  # which router to target; omit for the default
  group    = "read"
  name     = "tf-example"
  password = "REDACTED"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "replace-me"
  # alias = "replace-me"
  # inactivity_policy = "replace-me"
  # inactivity_timeout = "1h"
  # type = 0
}
