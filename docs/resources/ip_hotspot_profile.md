---
subcategory: "IP"
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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `dns_name` - (Optional) Type: `string`.
* `hotspot_address` - (Optional) Type: `ip`.
* `html_directory` - (Optional) Type: `string`.
* `html_directory_override` - (Optional) Type: `string`.
* `http_cookie_lifetime` - (Optional) Type: `duration`.
* `http_proxy` - (Optional) Type: `string`.
* `install_hotspot_queue` - (Optional) Type: `bool`.
* `login_by` - (Optional) Type: `list`.
* `name` - (Optional) Type: `string`.
* `smtp_server` - (Optional) Type: `ip`.
* `split_user_domain` - (Optional) Type: `bool`.
* `use_radius` - (Optional) Type: `bool`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `default` - Type: `bool`.

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
