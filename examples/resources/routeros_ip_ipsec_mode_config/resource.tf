resource "routeros_ip_ipsec_mode_config" "mode_config_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "example"

  # Optional attributes (uncomment as needed):
  # address = "10.99.0.1"
  # address_pool = "4.294967295e+09"
  # address_prefix_length = 24
  # connection_mark = "replace-me"
  # responder = false
  # split_dns = "replace-me"
  # split_include = "replace-me"
  # src_address_list = "my-list"
  # static_dns = "replace-me"
  # system_dns = false
  # use_responder_dns = "exclusively"
}
