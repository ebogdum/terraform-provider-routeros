resource "routeros_caps_man_provisioning" "provisioning_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "none"
  # common_name_regexp = "replace-me"
  # hw_supported_modes = "replace-me"
  # identity_regexp = "replace-me"
  # ip_address_ranges = "replace-me"
  # master_configuration = "replace-me"
  # name_format = "cap"
  # name_prefix = "replace-me"
  # radio_mac = "02:00:00:00:00:01"
  # slave_configurations = "replace-me"
}
