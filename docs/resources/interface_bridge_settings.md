---
subcategory: "Bridge"
page_title: "RouterOS: routeros_interface_bridge_settings"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_bridge_settings

Manages the RouterOS `/interface/bridge/settings` menu.

## Example Usage

```terraform
resource "routeros_interface_bridge_settings" "settings_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # allow_fast_path = false
  # use_ip_firewall = false
  # use_ip_firewall_for_pppoe = false
  # use_ip_firewall_for_vlan = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `allow_fast_path` - (Optional) Type: `bool`.
* `bridge_fast_forward_bytes` - (Read-only) Type: `int`.
* `bridge_fast_forward_packets` - (Read-only) Type: `int`.
* `bridge_fast_path_active` - (Read-only) Type: `bool`.
* `bridge_fast_path_bytes` - (Read-only) Type: `int`.
* `bridge_fast_path_packets` - (Read-only) Type: `int`.
* `use_ip_firewall` - (Optional) Type: `bool`.
* `use_ip_firewall_for_pppoe` - (Optional) Type: `bool`.
* `use_ip_firewall_for_vlan` - (Optional) Type: `bool`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_interface_bridge_settings.this 'home'
```
