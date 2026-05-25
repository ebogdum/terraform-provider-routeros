resource "routeros_special_login" "special_login_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # channel = "replace-me"
  # port = "443"
  # user = "myuser"
}
