resource "routeros_ip_smb" "smb_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # domain = "example.local"
  # enabled = "replace-me"
  # interface = "ether1"
  # interfaces = "replace-me"
}
