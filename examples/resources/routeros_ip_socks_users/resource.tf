resource "routeros_ip_socks_users" "users_example" {
  # router = "my-router"  # which router to target; omit for the default
  name     = "example"
  password = "REDACTED"

  disabled = false
}
