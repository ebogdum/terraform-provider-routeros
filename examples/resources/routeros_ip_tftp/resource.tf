resource "routeros_ip_tftp" "tftp_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # allow = true
  # ip_addresses = "replace-me"
  # read_only = true
  # real_filename = "replace-me"
  # req_filename = "replace-me"
}
