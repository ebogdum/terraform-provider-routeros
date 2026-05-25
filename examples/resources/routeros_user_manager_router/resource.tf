resource "routeros_user_manager_router" "router_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "replace-me"
  # coa_port = "443"
  # name = "tf-example"
  # protocol = "replace-me"
  # shared_secret = "REDACTED"
}
