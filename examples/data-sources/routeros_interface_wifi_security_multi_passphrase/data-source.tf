data "routeros_interface_wifi_security_multi_passphrase" "multi_passphrase_example" {
  # router = "my-router"  # omit for the default router
  # filter = { group = "guest-group" }

  # Omit proplist and every PSK lands in state in cleartext.
  proplist = [".id", "group", "vlan-id"]
}
