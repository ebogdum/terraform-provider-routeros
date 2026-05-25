resource "routeros_ip_kid_control" "kid_control_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # name = "tf-example"
  # rate_limit = "replace-me"
}
