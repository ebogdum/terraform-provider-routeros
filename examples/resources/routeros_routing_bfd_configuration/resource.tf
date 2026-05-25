resource "routeros_routing_bfd_configuration" "configuration_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address_list = "replace-me"
  # addresses = "replace-me"
  # forbid_bfd = "replace-me"
  # interfaces = "replace-me"
  # min_rx = "replace-me"
  # min_tx = "replace-me"
  # multiplier = "replace-me"
  # vrf = "main"
}
