resource "routeros_system_backup" "backup_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # dont_encrypt = "replace-me"
  # encryption = "replace-me"
  # name = "tf-example"
  # password = "REDACTED"
}
