resource "routeros_ip_hotspot" "hotspot_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  name      = "tf-example"

  disabled = false

  # Optional attributes (uncomment as needed):
  # dns_name = "replace-me"
  # hotspot_address = "10.99.0.0/24"
  # html_directory = "replace-me"
  # html_directory_override = "replace-me"
  # http_cookie_lifetime = "replace-me"
  # http_proxy = "replace-me"
  # https_redirect = "replace-me"
  # keepalive_timeout = "replace-me"
  # login_by = "replace-me"
  # mac_auth_password = "REDACTED"
  # nas_port_type = "replace-me"
  # profile = "replace-me"
  # radius_accounting = "replace-me"
  # radius_default_domain = "replace-me"
  # radius_interim_update = "replace-me"
  # radius_location_name = "replace-me"
  # radius_mac_format = "replace-me"
  # rate_limit = "replace-me"
  # smtp_server = "replace-me"
  # split_user_domain = "replace-me"
  # ssl_certificate = "replace-me"
  # trial_uptime = "replace-me"
  # trial_user_profile = "replace-me"
  # use_radius = "replace-me"
}
