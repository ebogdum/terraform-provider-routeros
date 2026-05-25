resource "routeros_ip_vrf" "vrf_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # interfaces = "replace-me"
  # name = "tf-example"
}
