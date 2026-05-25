resource "routeros_interface_vrrp" "vrrp_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  name      = "tf-example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # authentication = "replace-me"
  # interval = "replace-me"
  # password = "REDACTED"
  # priority = "replace-me"
  # remote_address = "10.99.0.1"
}
