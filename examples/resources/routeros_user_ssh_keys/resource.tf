resource "routeros_user_ssh_keys" "ssh_keys_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # import_ssh_key = "REDACTED"
  # key = "replace-me"
  # newk = "replace-me"
  # oldk = "replace-me"
}
