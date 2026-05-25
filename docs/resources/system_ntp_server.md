---
subcategory: "NTP"
page_title: "RouterOS: routeros_system_ntp_server"
description: |-
  RouterOS resource.
---

# Resource: routeros_system_ntp_server

Manages the RouterOS `/system/ntp/server` menu.

## Example Usage

```terraform
resource "routeros_system_ntp_server" "server_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # auth_key = "REDACTED"
  # broadcast = false
  # broadcast_addresses = "replace-me"
  # enabled = false
  # local_clock_stratum = 0
  # manycast = false
  # multicast = false
  # use_local_clock = false
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `auth_key` - (Optional) Type: `string`. **Sensitive.**
* `broadcast` - (Optional) Type: `bool`.
* `broadcast_addresses` - (Optional) Type: `string`.
* `enabled` - (Optional) Type: `bool`.
* `local_clock_stratum` - (Optional) Type: `int`.
* `manycast` - (Optional) Type: `bool`.
* `multicast` - (Optional) Type: `bool`.
* `use_local_clock` - (Optional) Type: `bool`.
* `vrf` - (Optional) Type: `string`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_ntp_server.this 'home'
```
