resource "routeros_ip_smb_users" "users_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # name = "example"
  # password = "REDACTED"
  # read_only = false
}
