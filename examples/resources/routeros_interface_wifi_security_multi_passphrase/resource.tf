resource "routeros_interface_wifi_security_multi_passphrase" "multi_passphrase_example" {
  # router = "my-router"  # which router to target; omit for the default
  group      = "guest-group"
  passphrase = "replace-me"

  # Optional attributes (uncomment as needed):
  # comment = "managed by terraform"
  # disabled = false
  # expires = "replace-me"
  # isolation = false
  # vlan_id = "replace-me"
}
