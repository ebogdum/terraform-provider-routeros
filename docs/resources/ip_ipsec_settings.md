---
subcategory: "IPsec"
page_title: "RouterOS: routeros_ip_ipsec_settings"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_ipsec_settings

Manages the RouterOS `/ip/ipsec/settings` menu.

## Example Usage

```terraform
resource "routeros_ip_ipsec_settings" "settings_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # accounting = false
  # ddos_cookie_threshold = 0
  # interim_update = "1h"
  # xauth_use_radius = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `accounting` - (Optional) Type: `bool`.
* `ddos_cookie_threshold` - (Optional) Type: `int`.
* `interim_update` - (Optional) Type: `duration`.
* `xauth_use_radius` - (Optional) Type: `bool`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_ipsec_settings.this 'home'
```
