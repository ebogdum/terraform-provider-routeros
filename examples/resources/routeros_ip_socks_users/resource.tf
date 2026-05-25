resource "routeros_ip_socks_users" "users_example" {
  # router = "my-router"  # which router to target; omit for the default
  name     = "tf-example"
  password = "REDACTED"

  disabled = false
}
