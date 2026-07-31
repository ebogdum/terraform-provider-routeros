---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_neighbor_discovery_settings"
description: |-
  Mirrors RouterOS /ip/neighbor/discovery-settings.
---

# Resource: routeros_ip_neighbor_discovery_settings

Mirrors RouterOS `/ip/neighbor/discovery-settings`.

## Example Usage

```terraform
resource "routeros_ip_neighbor_discovery_settings" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # discover_interface_list = "replace-me"
  # discover_interval = "replace-me"
  # lldp_mac_phy_config = true
  # lldp_max_frame_size = true
  # lldp_med_net_policy_vlan = "replace-me"
  # lldp_poe_power = true
  # lldp_vlan_info = true
  # mode = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `add_dns_entries` - (Optional) Type: `string`. RouterOS `add-dns-entries`.
* `add_dns_entries_suffix` - (Optional) Type: `string`. RouterOS `add-dns-entries-suffix`.
* `discover_interface_list` - (Optional) Type: `string`. RouterOS `discover-interface-list`.
* `discover_interval` - (Optional) Type: `string`. RouterOS `discover-interval`.
* `lldp_mac_phy_config` - (Optional) Type: `bool`. RouterOS `lldp-mac-phy-config`.
* `lldp_max_frame_size` - (Optional) Type: `bool`. RouterOS `lldp-max-frame-size`.
* `lldp_med` - (Optional) Type: `string`. RouterOS `lldp-med`.
* `lldp_med_net_policy_vlan` - (Optional) Type: `string`. RouterOS `lldp-med-net-policy-vlan`.
* `lldp_poe_power` - (Optional) Type: `bool`. RouterOS `lldp-poe-power`.
* `lldp_vlan_info` - (Optional) Type: `bool`. RouterOS `lldp-vlan-info`.
* `mode` - (Optional) Type: `string`. RouterOS `mode`.
* `protocol` - (Optional) Type: `string`. RouterOS `protocol`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_neighbor_discovery_settings.this 'home'
```
