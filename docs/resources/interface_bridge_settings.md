---
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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `allow_fast_path` - (Optional) Type: `bool`.
* `use_ip_firewall` - (Optional) Type: `bool`.
* `use_ip_firewall_for_pppoe` - (Optional) Type: `bool`.
* `use_ip_firewall_for_vlan` - (Optional) Type: `bool`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `bridge_fast_forward_bytes` - Type: `int`.
* `bridge_fast_forward_packets` - Type: `int`.
* `bridge_fast_path_active` - Type: `bool`.
* `bridge_fast_path_bytes` - Type: `int`.
* `bridge_fast_path_packets` - Type: `int`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_interface_bridge_settings.this 'home'
```
