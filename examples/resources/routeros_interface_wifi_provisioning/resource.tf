resource "routeros_interface_wifi_provisioning" "provisioning_example" {
  # router = "my-router"  # which router to target; omit for the default
  action = "create-dynamic-enabled"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address_ranges = "replace-me"
  # common_name_regexp = "replace-me"
  # identity_regexp = "replace-me"
  # master_configuration = "replace-me"
  # multi_link_mode = "replace-me"
  # name_format = "replace-me"
  # radio_mac = "replace-me"
  # slave_configurations = "replace-me"
  # slave_name_format = "replace-me"
  # supported_bands = "replace-me"
  # supported_hw_caps = "replace-me"
}
