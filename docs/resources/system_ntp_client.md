---
subcategory: "NTP"
page_title: "RouterOS: routeros_system_ntp_client"
description: |-
  RouterOS resource.
---

# Resource: routeros_system_ntp_client

Manages the RouterOS `/system/ntp/client` menu.

## Example Usage

```terraform
resource "routeros_system_ntp_client" "client_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # enabled = false
  # mode = "replace-me"
  # servers = "replace-me"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `enabled` - (Optional) Type: `bool`.
* `freq_drift` - (Optional) Type: `int`.
* `mode` - (Optional) Type: `string`.
* `servers` - (Optional) Type: `string`.
* `status` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_ntp_client.this 'home'
```
