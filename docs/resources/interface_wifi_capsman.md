---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_capsman"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_wifi_capsman

Manages the RouterOS `/interface/wifi/capsman` menu.

## Example Usage

```terraform
resource "routeros_interface_wifi_capsman" "capsman_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # enabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `ca_certificate` - (Optional) Type: `string`. RouterOS `ca-certificate`.
* `certificate` - (Optional) Type: `string`. RouterOS `certificate`.
* `enabled` - (Optional) Type: `bool`.
* `interfaces` - (Optional) Type: `string`. RouterOS `interfaces`.
* `package_path` - (Optional) Type: `string`. RouterOS `package-path`.
* `require_peer_certificate` - (Optional) Type: `string`. RouterOS `require-peer-certificate`.
* `upgrade_policy` - (Optional) Type: `string`. RouterOS `upgrade-policy`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_interface_wifi_capsman.this 'home'
```
