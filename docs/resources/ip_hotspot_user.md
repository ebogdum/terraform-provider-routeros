---
subcategory: "Hotspot"
page_title: "RouterOS: routeros_ip_hotspot_user"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_hotspot_user

Manages the RouterOS `/ip/hotspot/user` menu.

## Example Usage

```terraform
resource "routeros_ip_hotspot_user" "user_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "10.99.0.1"
  # email = "replace-me"
  # limit_bytes_in = "replace-me"
  # limit_bytes_out = "replace-me"
  # limit_bytes_total = "replace-me"
  # limit_uptime = "1h"
  # mac_address = "10.99.0.0/24"
  # otp_secret = "REDACTED"
  # password = "REDACTED"
  # profile = "replace-me"
  # routes = "replace-me"
  # server = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Optional) Type: `ip`.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `email` - (Optional) Type: `string`.
* `limit_bytes_in` - (Optional) Type: `string`.
* `limit_bytes_out` - (Optional) Type: `string`.
* `limit_bytes_total` - (Optional) Type: `string`.
* `limit_uptime` - (Optional) Type: `duration`.
* `mac_address` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_user`.
* `otp_secret` - (Optional) Type: `string`. **Sensitive.**
* `password` - (Optional) Type: `string`. **Sensitive.**
* `profile` - (Optional) Type: `string`.
* `routes` - (Optional) Type: `string`.
* `server` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `bytes_in` - Type: `int`.
* `bytes_out` - Type: `int`.
* `default` - Type: `bool`.
* `dynamic` - Type: `bool`.
* `packets_in` - Type: `int`.
* `packets_out` - Type: `int`.
* `uptime` - Type: `duration`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_hotspot_user.example '*3'

# Named router
terraform import routeros_ip_hotspot_user.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_hotspot_user.example 'home/my-resource-name'
```
