resource "routeros_routing_rip_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # cost = "replace-me"
  # instance = "replace-me"
  # interfaces = "replace-me"
  # key_chain = "replace-me"
  # mode = "replace-me"
  # name = "example"
  # password = "REDACTED"
  # poison_reverse = "replace-me"
  # source_addresses = "replace-me"
  # split_horizon = "replace-me"
  # use_bfd = "replace-me"
}
