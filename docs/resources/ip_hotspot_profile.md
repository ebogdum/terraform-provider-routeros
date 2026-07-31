---
subcategory: "Hotspot"
page_title: "RouterOS: routeros_ip_hotspot_profile"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_hotspot_profile

Manages the RouterOS `/ip/hotspot/profile` menu.

## Example Usage

```terraform
resource "routeros_ip_hotspot_profile" "profile_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # dns_name = "replace-me"
  # hotspot_address = "10.99.0.0/24"
  # html_directory = "replace-me"
  # html_directory_override = "replace-me"
  # http_cookie_lifetime = "1h"
  # http_proxy = "replace-me"
  # install_hotspot_queue = false
  # login_by = []
  # name = "tf-example"
  # smtp_server = "10.99.0.1"
  # split_user_domain = false
  # use_radius = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `default` - (Read-only) Type: `bool`.
* `dns_name` - (Optional) Type: `string`.
* `hotspot_address` - (Optional) Type: `string`.
* `html_directory` - (Optional) Type: `string`.
* `html_directory_override` - (Optional) Type: `string`.
* `http_cookie_lifetime` - (Optional) Type: `string`.
* `http_proxy` - (Optional) Type: `string`.
* `install_hotspot_queue` - (Optional) Type: `bool`.
* `login_by` - (Optional) Type: `list`.
* `mac_auth_mode` - (Optional) Type: `string`. RouterOS `mac-auth-mode`.
* `mac_auth_password` - (Optional) Type: `string`. RouterOS `mac-auth-password`. **Sensitive.**
* `name` - (Optional) Type: `string`.
* `nas_port_type` - (Optional) Type: `string`. RouterOS `nas-port-type`.
* `radius_accounting` - (Optional) Type: `string`. RouterOS `radius-accounting`.
* `radius_default_domain` - (Optional) Type: `string`. RouterOS `radius-default-domain`.
* `radius_interim_update` - (Optional) Type: `string`. RouterOS `radius-interim-update`.
* `radius_location_id` - (Optional) Type: `string`. RouterOS `radius-location-id`.
* `radius_location_name` - (Optional) Type: `string`. RouterOS `radius-location-name`.
* `radius_mac_format` - (Optional) Type: `string`. RouterOS `radius-mac-format`.
* `rate_limit` - (Optional) Type: `string`. RouterOS `rate-limit`.
* `smtp_server` - (Optional) Type: `string`.
* `split_user_domain` - (Optional) Type: `bool`.
* `ssl_certificate` - (Optional) Type: `string`. RouterOS `ssl-certificate`.
* `trial_uptime_limit` - (Optional) Type: `string`. RouterOS `trial-uptime-limit`.
* `trial_uptime_reset` - (Optional) Type: `string`. RouterOS `trial-uptime-reset`.
* `trial_user_profile` - (Optional) Type: `string`. RouterOS `trial-user-profile`.
* `use_radius` - (Optional) Type: `bool`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_hotspot_profile.example '*3'

# Named router
terraform import routeros_ip_hotspot_profile.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_hotspot_profile.example 'home/my-resource-name'
```
