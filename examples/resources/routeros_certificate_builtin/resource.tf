resource "routeros_certificate_builtin" "builtin_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # akid = "replace-me"
  # common_name = "replace-me"
  # country = "replace-me"
  # days_valid = 0
  # invalid_after = "replace-me"
  # invalid_before = "replace-me"
  # issuer = "replace-me"
  # key_type = "replace-me"
  # key_usage = []
  # locality = "replace-me"
  # organization = "replace-me"
  # serial_number = "replace-me"
  # skid = "replace-me"
  # state = "replace-me"
  # subject_alt_name = "replace-me"
  # unit = "replace-me"
}
