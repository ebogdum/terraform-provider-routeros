resource "routeros_ip_smb_shares" "shares_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # directory = "replace-me"
  # invalid_users = "replace-me"
  # read_only = false
  # require_encryption = false
  # valid_users = "replace-me"
}
