resource "routeros_caps_man_security" "security_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "example"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # encryption = "replace-me"
  # passphrase = "replace-me"
}
