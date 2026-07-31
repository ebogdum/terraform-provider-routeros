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
  # def = false
  # email = "replace-me"
  # limit_bytes_in = "replace-me"
  # limit_bytes_out = "replace-me"
  # limit_bytes_total = "replace-me"
  # limit_uptime = "1h"
  # mac_address = "10.99.0.0/24"
  # nondef = "replace-me"
  # nondefro = "replace-me"
  # otp_secret = "REDACTED"
  # password = "REDACTED"
  # profile = "replace-me"
  # reset_all_counters = "replace-me"
  # reset_counters = "replace-me"
  # routes = "replace-me"
  # server = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address` - (Optional) Type: `string`.
* `bytes_in` - (Read-only) Type: `int`.
* `bytes_out` - (Read-only) Type: `int`.
* `comment` - (Optional) Type: `string`.
* `def` - (Optional) Type: `bool`.
* `default` - (Read-only) Type: `bool`.
* `disabled` - (Optional) Type: `bool`.
* `dynamic` - (Read-only) Type: `bool`.
* `email` - (Optional) Type: `string`.
* `limit_bytes_in` - (Optional) Type: `string`.
* `limit_bytes_out` - (Optional) Type: `string`.
* `limit_bytes_total` - (Optional) Type: `string`.
* `limit_uptime` - (Optional) Type: `string`.
* `mac_address` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `nondef` - (Optional) Type: `string`.
* `nondefro` - (Read-only) Type: `string`.
* `otp_secret` - (Optional) Type: `string`. **Sensitive.**
* `packets_in` - (Read-only) Type: `int`.
* `packets_out` - (Read-only) Type: `int`.
* `password` - (Optional) Type: `string`. **Sensitive.**
* `profile` - (Optional) Type: `string`.
* `reset_all_counters` - (Read-only) Type: `string`.
* `reset_counters` - (Read-only) Type: `string`.
* `routes` - (Optional) Type: `string`.
* `server` - (Optional) Type: `string`.
* `uptime` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
