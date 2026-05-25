---
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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `add_mac_cookie` - (Optional) Type: `bool`.
* `address_list` - (Optional) Type: `string`.
* `idle_timeout` - (Optional) Type: `string`.
* `keepalive_timeout` - (Optional) Type: `duration`.
* `mac_cookie_timeout` - (Optional) Type: `duration`.
* `name` - (Optional) Type: `string`.
* `shared_users` - (Optional) Type: `int`.
* `status_autorefresh` - (Optional) Type: `duration`.
* `transparent_proxy` - (Optional) Type: `bool`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `default` - Type: `bool`.

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
