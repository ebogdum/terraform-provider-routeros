---
subcategory: "Hotspot"
page_title: "RouterOS: routeros_ip_hotspot_user_profile"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_hotspot_user_profile

Manages the RouterOS `/ip/hotspot/user/profile` menu.

## Example Usage

```terraform
resource "routeros_ip_hotspot_user_profile" "profile_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # add_mac_cookie = false
  # address_list = "replace-me"
  # idle_timeout = "replace-me"
  # keepalive_timeout = "1h"
  # mac_cookie_timeout = "1h"
  # name = "tf-example"
  # shared_users = 0
  # status_autorefresh = "1h"
  # transparent_proxy = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `add_mac_cookie` - (Optional) Type: `bool`.
* `address_list` - (Optional) Type: `string`.
* `default` - (Read-only) Type: `bool`.
* `idle_timeout` - (Optional) Type: `string`.
* `keepalive_timeout` - (Optional) Type: `string`.
* `mac_cookie_timeout` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `shared_users` - (Optional) Type: `string`. A number, or `unlimited`.
* `status_autorefresh` - (Optional) Type: `string`.
* `transparent_proxy` - (Optional) Type: `bool`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_hotspot_user_profile.example '*3'

# Named router
terraform import routeros_ip_hotspot_user_profile.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_hotspot_user_profile.example 'home/my-resource-name'
```
